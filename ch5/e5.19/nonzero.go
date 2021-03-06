package main

import (
	"fmt"
)

func nonzero(n int) (r int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("no panic?")
		case "zero":
			r = 1
		case "not zero":
			r = n
		default:
			fmt.Println("I don't know who called me")
			panic(p)
		}
	}()
	if n == 0 {
		panic("zero")
	} else {
		panic("not zero")
	}
}

func main() {
	fmt.Println(nonzero(0))
	fmt.Println(nonzero(10))
}
