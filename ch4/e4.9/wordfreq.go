package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	words := make(map[string]int)
	for input.Scan() {
		words[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for w, n := range words {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
