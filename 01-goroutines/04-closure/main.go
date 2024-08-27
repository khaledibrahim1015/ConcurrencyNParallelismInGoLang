package main

import (
	"fmt"
	"sync"
)

func main() {

	//  run the program and check that variable i
	// was pinned for access from goroutine even after
	// enclosing function returns.
	// output
	// return from function
	// value of i: 1
	// done..
	var wg sync.WaitGroup

	// anonymous function
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Println("return from function")
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")

	// another impl with outpput
	var wg2 sync.WaitGroup

	incr2 := func(wg *sync.WaitGroup) {

		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++

			fmt.Printf("value of i: %v\n", i)
		}()
		wg.Wait()
		fmt.Println("return from function")
		return

	}

	incr2(&wg2)
	wg.Wait()
	fmt.Println("done..")

}
