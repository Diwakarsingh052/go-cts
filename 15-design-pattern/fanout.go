package main

import (
	"fmt"
	"strconv"
)

func main() {

	const job = 10
	ch := make(chan string, job)
	// for each task we are running one go routine
	for work := 1; work <= 10000; work++ {
		//always copy value to the go routine param. we should not access it directly // remove the param to see the side effects
		go func(w int) {
			// do task here
			ch <- "work" + strconv.Itoa(w)
		}(work)

	}

	for i := 1; i <= 10000; i++ {
		fmt.Println(<-ch)
	}
}

/*
	for range url slice {
	go func() {http.get}
*/
