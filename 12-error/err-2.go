package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not able to find in root directory")

func main() {
	f, err := openFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
}

func openFile(fileName string) (*os.File, error) {

	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		//errorf returns an error value // it is same as printf, but it returns the values instead of printing it
		//before go 1.13 approach // %v is used to merge the value
		return nil, fmt.Errorf("custom msg %v %v", err, ErrFileNotFound) // when err happens then set other values to their defaults
	}
	return f, nil

}
