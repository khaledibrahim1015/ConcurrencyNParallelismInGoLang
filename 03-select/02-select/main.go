package main

import (
	"fmt"
	"time"
)

func main() {
	// impelemnt Timeout for recv on channel ch

	ch := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "one"
	}()

	select {
	case m1 := <-ch:
		fmt.Println(m1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")

	}
}
