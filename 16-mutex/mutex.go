package main

import (
	"fmt"
	"sync"
	"time"
)

var cab = 2
var wg = &sync.WaitGroup{}

func main() {
	m := &sync.Mutex{}
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		wg.Add(1)
		go bookCab(name, m)
	}
	wg.Wait()
	fmt.Println(cab)

}

func bookCab(name string, m *sync.Mutex) {
	defer wg.Done()

	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)

	// critical section where we are accessing the shared resource which is our global var cabs
	m.Lock()
	if cab >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("car is not available for", name)
	}
	m.Unlock()

}
