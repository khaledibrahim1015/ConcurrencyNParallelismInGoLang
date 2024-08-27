package main

import "fmt"

func main() {

	// recive  only channel
	owner := func() <-chan int {

		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		//  read values from channel
		for value := range ch {
			fmt.Println(value)
		}
	}

	ch := owner()
	consumer(ch)

}
