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
	mustCopy(zone, os.Stdout, conn)
	//mustCopy(conn, os.Stdin)
}

func mustCopy(zone string, dst io.Writer, src io.Reader) {
	dst.Write([]byte(zone + ": "))
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/*
go run ch8\e8.1\clockwall.go NewYork=192.168.88.108:8010 Tokyo=192.168.88.108:8020 London=192.168.88.108:8030
NewYork: Sun Oct 14 05:45:36 EDT 2018
Tokyo: Sun Oct 14 18:45:36 JST 2018
London: Sun Oct 14 10:45:36 BST 2018
*/
