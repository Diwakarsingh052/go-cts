package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

// we can make function calls concurrent
func main() {
	wg.Add(4)
	go hello() // func call // the go keyword spins up a new goroutine //
	go hello() // each func call creates a go routine // separate line exec // new goroutine // not care the order of exec
	go hello()
	go hello()
	fmt.Println("hello from the main")

	// time.sleep is not a good idea in production
	//time.Sleep(3 * time.Second) // sleeping is a non productive cpu activity so cpu will make the switch to the another go routine
	wg.Wait()
}

func hello() {
	defer wg.Done()
	time.Sleep(4 * time.Second) // to show latency only // not to be used in the prod code
	fmt.Println("hello from the hello func")
}
