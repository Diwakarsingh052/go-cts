package main

import (
	"fmt"
	"time"
)

//var wg = &sync.WaitGroup{}

func main() {
	c := make(chan int) // unbuffered chan
	//wg.Add(4)
	go addNum(10, 20, c)
	go sub(10, 5, c)
	go mutl(2, 2, c)
	//calcAll(c)
	//wg.Wait()
	x, y, z := <-c, <-c, <-c // recv from the channel // it is a blocking call // this line will block until we recv all the values
	fmt.Println(x, y, z)
}

func addNum(a, b int, c chan int) {
	//defer wg.Done()
	sum := a + b
	// unbuffered channel gives you a guarantee that the work would be always received by someone
	c <- sum // send over a channel // signaling with data // this line will block until a receiver is not ready
}

func sub(a, b int, c chan int) {
	//defer wg.Done()
	time.Sleep(1 * time.Second)
	sum := a - b
	c <- sum // send over a channel // signaling with data

}

func mutl(a, b int, c chan int) {
	//defer wg.Done()
	sum := a * b
	c <- sum // send over a channel // signaling with data

}

//func calcAll(c chan int) {
//	//defer wg.Done()
//	x, y, z := <-c, <-c, <-c // recv from the channel // it is a blocking call // this line will block until we recv all the values
//	fmt.Println(x, y, z)
//}
