package main

import "fmt"

func main() {

	ch := make(chan int)
	go func(x int, y int) {
		c := x + y
		ch <- c

	}(2, 3)

	// get value computed fom oroutine
	r := <-ch
	fmt.Printf("comuted value %v \n ", r)

}
