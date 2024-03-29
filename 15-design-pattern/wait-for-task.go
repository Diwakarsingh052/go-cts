package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {
	waitForTask()
}

type response struct {
	url  string
	resp *http.Response
	err  error
}

func waitForTask() {
	//u := ""
	wg.Add(2)
	url := make(chan string)
	respCh := make(chan response)

	go func() { //(url string)
		defer wg.Done()
		task := <-url // unknown latency // recv
		res, err := http.Get(task)
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//b, _ := io.ReadAll(res.Body)
		//fmt.Println(string(b))
		r := response{
			url:  task,
			resp: res,
			err:  err,
		}
		respCh <- r //sending the resp struct to respCh

	}()

	go func() {
		// process the result of the get request // not recommended // just to show how the communication works
		defer wg.Done()
		r := <-respCh

		if r.err != nil {
			log.Println(r.err)
			return
		}

		b, err := io.ReadAll(r.resp.Body)

		if err != nil {
			log.Println(r.err)
			return
		}
		defer r.resp.Body.Close()
		fmt.Println(string(b))
		fmt.Println(r.resp.Status)

	}()
	time.Sleep(2 * time.Second)
	url <- "https://pkg.go.dev/" //send
	wg.Wait()
}
