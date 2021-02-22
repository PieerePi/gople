// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

//!+

var shaType = flag.String("sha", "256", `"224" "256" "384" "512"`)

func main() {
	flag.Parse()

	switch *shaType {
	case "224", "256", "384", "512":
	default:
		fmt.Printf("Invalid shasum type: %s\n", *shaType)
		flag.PrintDefaults()
		os.Exit(1)
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		userInput := []byte(input.Text())
		switch *shaType {
		case "224":
			shasum := sha256.Sum224(userInput)
			fmt.Printf("%x\n", shasum)
		case "256":
			shasum := sha256.Sum256(userInput)
			fmt.Printf("%x\n", shasum)
		case "384":
			shasum := sha512.Sum384(userInput)
			fmt.Printf("%x\n", shasum)
		case "512":
			shasum := sha512.Sum512(userInput)
			fmt.Printf("%x\n", shasum)
		}
	}
}

//!-
