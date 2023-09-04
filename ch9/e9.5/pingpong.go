package main

import (
	"fmt"
	"os"
	"time"
)

var counter int64

func pingpong(in <-chan interface{}, out chan<- interface{}) {
	for v := range in {
		counter++
		out <- v
	}
}

func main() {
	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	in := make(chan interface{})
	out := make(chan interface{})
	go func() { pingpong(in, out) }()
	go func() { pingpong(out, in) }()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var lastCounter int64

	in <- interface{}(0)
	start := time.Now()

	for {
		select {
		case <-ticker.C:
			curCounter := counter
			fmt.Println(curCounter-lastCounter, "communications in last second.")
			lastCounter = curCounter
		case <-done:
			fmt.Printf("pingpong average throughput = %d communications per second.\n",
				counter*1000000000/time.Since(start).Nanoseconds())
			return
		}
	}
}
