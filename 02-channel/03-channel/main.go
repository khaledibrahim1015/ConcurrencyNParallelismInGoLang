package main

import (
	"fmt"
)

func main() {
	// ch := make(chan int) //  unbuffered chanel that mean it is blocking

	//  to make it unblocking make it buffer  channel
	ch := make(chan int, 6)

	go func() {
		defer close(ch)

		//  send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
