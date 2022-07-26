package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	//wg.Add(10) // counter to add number of goroutines that we are starting or spinning up
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go work(i, wg) // 10 goroutines
	}
	wg.Add(2)
	go work(11, wg)
	go calc(10, 30, wg) // from goroutines, we can't return values // so we create a blanket func and do our actual work there
	wg.Wait()           // main will wait until counter resets to zero
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the counter by one
	fmt.Println("i am working on ", i)
}

func calc(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	s := add(a, b)
	fmt.Println("sum ", s)
	a, err := strconv.Atoi("abc")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(a)

}

func add(a, b int) int {
	return a + b
}
