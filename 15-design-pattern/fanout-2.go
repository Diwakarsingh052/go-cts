package main

import (
	"fmt"
	"strconv"
)

func main() {
	const work = 10
	const cap = 2
	ch := make(chan string, work)
	sem := make(chan bool, cap)

	//task of this goroutine is to do fan out of the jobs
	go func() {

		for e := 1; e <= 100; e++ {
			go func(e int) {
				sem <- true // this line will block if there is no room in the buffer // otherwise we can proceed to the next line
				ch <- "task" + strconv.Itoa(e)
			}(e)
			//time.Sleep(1 * time.Second)
			<-sem // it will take the one value out of the buffer which will make the space for the other goroutine
		}
		ch <- "101"
		close(ch) // close channels from where you are sending the data // panic: send on closed channel

	}()
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(<-ch)
	//}
	go func() {
		for v := range ch { // recv over the channel until the channel is not closed // we can recv over the closed channel but we can't send more data to the closed one that is a panic situation
			fmt.Println(v)
		}
	}()

	abc := make(chan int)
	<-abc

}
