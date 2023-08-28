// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

//!+
func main() {
	for _, v := range os.Args[1:] {
		sl := strings.Split(v, "=")
		// no go func
		getTime(sl[0], sl[1])
	}
}

//!-

func getTime(zone string, ts string) {
	conn, err := net.Dial("tcp", ts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	os.Stdout.Write([]byte(zone + ": "))
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/*
go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
NewYork: Tue Mar  9 23:17:59 CST 2021
Tokyo: Tue Mar  9 23:18:00 CST 2021
London: Tue Mar  9 23:18:00 CST 2021
*/
