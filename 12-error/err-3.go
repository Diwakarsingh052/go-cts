package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("not able to find in root directory")

func main() {
	_, err := openFile("test.txt")

	if errors.Is(err, ErrFileNotFound) {
		fmt.Println("it is inside the chain")
	} else {
		fmt.Println("it is not in the chain")
	}

}

func openFile(fileName string) (*os.File, error) {

	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		//errorf returns an error value // it is same as printf, but it returns the values instead of printing it
		//%w means wrapping of the error
		return nil, fmt.Errorf("%v %w", err, ErrFileNotFound) // when err happens then set other values to their defaults
	}
	return f, nil

}
