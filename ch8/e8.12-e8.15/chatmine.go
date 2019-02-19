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

type msg struct {
	msgtype int           // 0 - enter, 1 - leave, 2 - normal
	sender  string        // sender
	message string        // 0 - addr, 1 - empty, 2 - message
	channel chan<- string // 0 - not nil, other - nil
}

type client struct {
	channel  chan<- string // an outgoing message channel
	addr     string
	name     string
	lasttime time.Time
}

var messages = make(chan msg) // all incoming client messages
var checkAlarm = make(chan bool)

func checkerAlarm() {
	for {
		time.Sleep(time.Second)
		checkAlarm <- true
	}
}

func broadcaster() {
	clients := make(map[string]client) // all connected clients
	for {
		select {
		case msg := <-messages:
			switch msg.msgtype {
			case 0:
				for _, cli := range clients {
					cli.channel <- msg.sender + " 进入了聊天室。"
				}
				clients[msg.sender] = client{msg.channel, msg.message, msg.sender, time.Now()}
			case 1:
				if _, ok := clients[msg.sender]; !ok {
					continue
				}
				close(clients[msg.sender].channel)
				delete(clients, msg.sender)
				for _, cli := range clients {
					cli.channel <- msg.sender + " 离开了聊天室。"
				}
			case 2:
				// Broadcast incoming message to all
				// clients' outgoing message channels.
				if _, ok := clients[msg.sender]; !ok {
					continue
				}
				clients[msg.sender] = client{clients[msg.sender].channel, clients[msg.sender].addr, clients[msg.sender].name, time.Now()}
				for _, cli := range clients {
					if !(msg.sender == cli.name) {
						cli.channel <- "\t\t" + msg.sender + " 说: " + msg.message
					}
				}
			}
		case _ = <-checkAlarm:
			for _, cli := range clients {
				if cli.lasttime.Add(30 * time.Second).Before(time.Now()) {
					//fmt.Println(cli.name + " 超时。")
					close(clients[cli.name].channel)
					delete(clients, cli.name)
					for _, cli2 := range clients {
						cli2.channel <- cli2.name + " 长时间没有反应，已经离开了聊天室。"
					}
				}
			}
		}
	}
}

func handleConn(conn net.Conn) {
	var who string
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	fmt.Fprintln(conn, "请输入你的名称:")
	input := bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	} else {
		conn.Close()
		return
	}
	//who := conn.RemoteAddr().String()
	ch <- who + "你好！"
	//messages <- msg{2, who, who + " has arrived", nil}
	messages <- msg{0, who, conn.RemoteAddr().String(), ch}

	//input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- msg{2, who, input.Text(), nil}
	}
	// NOTE: ignoring potential errors from input.Err()

	messages <- msg{1, who, "", nil}
	//messages <- msg{2, who, who + " has left", nil}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
	//fmt.Println("close conn")
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	go checkerAlarm()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
