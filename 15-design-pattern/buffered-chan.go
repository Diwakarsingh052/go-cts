package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

//A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	wg.Add(1)
	ch := make(chan int, 10)

	go func() {
		defer wg.Done()
		ch <- 1000
	}()
	fmt.Println(<-ch)
	wg.Wait()
}
