package main

import (
	"fmt"
	"sync"
)

//https://go.dev/ref/spec#Send_statements
//A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
//A send on a buffered channel can proceed if there is room in the buffer.
// receiver will always block until it is not able to receive values
var wg = &sync.WaitGroup{}

func main() {

	ch := make(chan int) // unbuffered chan
	a := 20
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- a // send on an unbuffered channel
	}()

	go func() {
		defer wg.Done()
		x := <-ch // recv from the channel // it is a blocking call // it will block until we are not able to fetch any values
		fmt.Println(x)
	}()

	wg.Wait()
}
