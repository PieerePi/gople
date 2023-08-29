// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
const (
	ENTERING  = 0
	LEAVING   = 1
	MESSAGING = 2
	TIMEOUT   = 3
)

const IDLETIMEOUT = 5 * time.Minute

type clientChannel chan<- string

type msg struct {
	msgtype int           // ENTERING, LEAVING, MESSAGING, TIMEOUT
	channel clientChannel // client outgoing message channel
	message string        // ENTERING - addr, MESSAGING - message, other - empty
	sender  string        // ENTERING - sender name, other - empty
}

type client struct {
	channel  clientChannel // an outgoing message channel
	addr     string
	name     string
	lasttime time.Time
}

var (
	clients  = make(map[clientChannel]client) // all connected clients
	messages = make(chan msg)                 // all incoming client messages
)

func writeOrSkip(channel clientChannel, message string) {
	select {
	case channel <- message:
	default:
	}
}

func broadcaster() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case msg := <-messages:
			switch msg.msgtype {
			case ENTERING:
				message := "\t\t" + msg.sender + " 进入了聊天室。"
				for channel := range clients {
					writeOrSkip(channel, message)
				}
				clients[msg.channel] = client{msg.channel, msg.message, msg.sender, time.Now()}
			case LEAVING, TIMEOUT:
				var sender string
				var message string
				if _, ok := clients[msg.channel]; !ok {
					// TIMEOUT: we will get a LEAVING here
					continue
				} else {
					sender = clients[msg.channel].name
					if msg.msgtype == TIMEOUT {
						writeOrSkip(msg.channel, sender+" 由于你长时间没有反应，你已被移出了聊天室。")
					}
					close(msg.channel)
					delete(clients, msg.channel)
				}
				if msg.msgtype == LEAVING {
					message = "\t\t" + sender + " 离开了聊天室。"
				} else {
					message = "\t\t" + sender + " 长时间没有反应，已经离开了聊天室。"
				}
				for channel := range clients {
					writeOrSkip(channel, message)
				}
			case MESSAGING:
				// Broadcast incoming message to all
				// clients' outgoing message channels.
				var sender string
				if _, ok := clients[msg.channel]; !ok {
					continue
				} else {
					sender = clients[msg.channel].name
					// Update lasttime
					orgClient := clients[msg.channel]
					orgClient.lasttime = time.Now()
					clients[msg.channel] = orgClient
				}
				message := "\t\t" + sender + " 说: " + msg.message
				for channel := range clients {
					if channel != msg.channel {
						writeOrSkip(channel, message)
					}
				}
			}
		case <-ticker.C:
			for _, cli := range clients {
				if cli.lasttime.Add(IDLETIMEOUT).Before(time.Now()) {
					lch := cli.channel
					go func() { messages <- msg{TIMEOUT, lch, "", ""} }()
				}
			}
		}
	}
}

func handleConn(conn net.Conn) {
	var who string
	ch := make(chan string, 5) // outgoing client messages, buffered channel
	go clientWriter(conn, ch)

	fmt.Fprint(conn, "请输入你的名称：")
	input := bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	} else {
		conn.Close()
		return
	}
	ch <- who + " 你好！"
	messages <- msg{ENTERING, ch, conn.RemoteAddr().String(), who}

	for input.Scan() {
		messages <- msg{MESSAGING, ch, input.Text(), ""}
	}
	// NOTE: ignoring potential errors from input.Err()

	messages <- msg{LEAVING, ch, "", ""}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
	// TIMEOUT: close the connection to let handleConn exit the input.Scan loop
	// LEAVING: the connection might be closed in handleConn first or here?
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
