package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(3)
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		defer wg.Done()

		c1 <- "one" // send
	}()
	go func() {
		defer wg.Done()

		c2 <- "two"
	}()

	go func() {
		defer wg.Done()

		c3 <- "three"
	}()
	for i := 1; i <= 3; i++ {
		// whichever case is not blocking exec that first
		//whichever case is ready first exec that.
		select {
		case x := <-c1: // if this chan is ready with values exec the case
			fmt.Println(x)
		case y := <-c2:
			fmt.Println(y)
		case z := <-c3:
			fmt.Println(z)

		}
	}
	//fmt.Println(<-c1) // recv
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)
	wg.Wait()
}
