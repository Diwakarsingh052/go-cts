package main

import (
	"fmt"
	"time"
)

func main() {

}

func waitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		ch <- "result"
	}()

	r := <-ch // unknown latency
	//some other tasks that would be blocked until the recv happens
	fmt.Println(r)
}
