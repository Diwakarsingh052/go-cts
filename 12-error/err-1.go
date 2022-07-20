package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

var ErrRecordNotFound = errors.New("record not found with the id provided")

func fetchRecord(id int) (string, error) {

	name, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}
	return name, nil

}

func main() {
	data, err := fetchRecord(100)
	//var s string = err.Error() // convert err to string
	//fmt.Println(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
