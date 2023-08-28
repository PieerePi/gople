// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// !+
func handleConn(c net.Conn) {
	reportChan := make(chan int)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			// got text
			reportChan <- 1
			go echo(c, input.Text(), 1*time.Second)
		}
		// NOTE: ignoring potential errors from input.Err()
		c.Close()
		// closed
		reportChan <- 0
	}()

	gotRequest := false
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if !gotRequest {
				log.Println("No request in 10 seconds, close the connection.")
				// NOTE: ignoring potential errors from input.Err()
				c.Close()
			} else {
				gotRequest = false
			}
		case r := <-reportChan:
			if r == 0 {
				log.Println("Connection closed.")
				return
			} else if r == 1 {
				gotRequest = true
			}
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
