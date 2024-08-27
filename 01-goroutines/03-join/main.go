package main

import (
	"fmt"
	"sync"
)

func main() {
	//modify the program
	// to print the value as 1
	// deterministically.
	var data int

	var wg sync.WaitGroup
	wg.Add(1)
	//  make it awaitable
	go func() {
		defer wg.Done() //  decrement number of specified goroutines
		data++
	}()
	//  to block main rouine (thread )
	wg.Wait() //  unblock when become 0

	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
