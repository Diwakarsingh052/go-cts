package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type cabs struct {
	driver int
	rw     sync.RWMutex
	wg     sync.WaitGroup
}

func (c *cabs) bookCab(name string) {
	defer c.wg.Done()
	c.rw.Lock()
	fmt.Println("welcome to the website", name)
	//c.rw.Lock() // write lock // no one can read if a goroutine is writing //only one goroutine can enter to write.
	if c.driver >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		c.driver--
	} else {
		fmt.Println("car is not available for", name)
	}
	c.rw.Unlock() // write unlock
}

func (c *cabs) getCabDriver() {
	defer c.wg.Done()
	c.rw.RLock() // read lock:when a goroutine is reading then no other can be writing to the shared resource// there could be unlimited number of reads
	fmt.Println("driver", c.driver)
	c.rw.RUnlock()

}

func main() {
	c := cabs{
		driver: 5,
	}

	for i := 1; i <= 15; i++ {
		c.wg.Add(1)
		go c.getCabDriver()
	}
	for i := 1; i <= 15; i++ {
		c.wg.Add(1)
		go c.bookCab("user " + strconv.Itoa(i))

	}

	for i := 1; i <= 15; i++ {
		c.wg.Add(1)
		go c.getCabDriver()
	}
	//log.New()

	c.wg.Wait()
}
