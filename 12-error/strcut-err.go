package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNotFound = errors.New("not found")
var ErrMismatch = errors.New("mismatch")

var student = map[int]string{
	1: "raj",
}

type QueryError struct {
	Func  string
	Input string
	Err   error
}

// this method is compulsory to format your err msg // and all the struct field should be used inside the error msg
func (q *QueryError) Error() string {
	return "main." + q.Func + ": parsing " + q.Input + " " + q.Err.Error()
}
func (q *QueryError) Unwrap() error { return q.Err }

func SearchSomething(id int) error {

	name, ok := student[id]
	if !ok {
		return &QueryError{
			Func:  "SearchSomething",
			Input: strconv.Itoa(id),
			Err:   ErrNotFound,
		}
	}
	fmt.Println(name)
	return nil

}

func SomeOtherFunc(data string) error {
	return &QueryError{
		Func:  "SomeOtherFunc",
		Input: data,
		Err:   ErrMismatch,
	}
}

func main() {
	//a := 10
	//a = 20
	//a = 30
	//fmt.Println(a)

	err := SearchSomething(10)
	fmt.Println(err)
	err = SomeOtherFunc("hello")
	//fmt.Println(err)
	//_, err = strconv.Atoi("abc")
	//os.PathError{}

	var e *QueryError // nil // pointer is compulsory // value would be nil
	//var n *strconv.NumError

	if errors.As(err, &e) { // reference imp // it checks whether struct is inside the error chain or not
		fmt.Println("it is present in the chain", e.Func)
	} else {
		fmt.Println("not")
	}

	//Unwrap returns nil if the unwrap method is not implemented over the custom struct
	fmt.Println(errors.Unwrap(err))

}
