package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type abc struct{}

func (a abc) Read(p []byte) (n int, err error) {

	return 0, nil
}

func main() {
	var ErrNotFound = errors.New("any err msg")
	fmt.Println(ErrNotFound)

	//io.Reader()
	res, err := http.Get("https://go.dev/")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
