package main

import (
	"errors"
	"fmt"
)

var ErrFeesNotSubmitted error = errors.New("fees not submitted")
var ErrAdmissionCancelled error = errors.New("admission cancelled")
var ErrDetails = errors.New("wrong details provided")

func admission() error {
	err := fees()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrAdmissionCancelled)
	}
	return nil
}

func fees() error {
	err := details()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrFeesNotSubmitted)
	}
	return nil
}

func details() error {
	return fmt.Errorf("%w", ErrDetails)
}

func main() {
	err := admission()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var ErrNotFound = errors.New("not found")
	//var ErrNotThere = errors.New("not found")
	//var ErrNot = errors.New("not found")
	//err := fmt.Errorf("%w %w %w", ErrNot, ErrNotThere, ErrNotFound) // we can only wrap one error at a time
	//fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)

}
