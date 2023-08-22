package main

import (
	"fmt"
	"os"

	"github.com/PieerePi/gople/ch7/e7.18/xmltree"
)

func main() {
	e, err := xmltree.Build(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xmltree: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", e)
}
