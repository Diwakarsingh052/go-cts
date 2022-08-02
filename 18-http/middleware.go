package main

import (
	"fmt"
	"log"
	"net/http"
)

// normal flow//  req -> handlerFunc -> FuncCall
// middleware flow// req -> mid1 -> mid2 -> handlerFunc ->  FuncCall
func main() {
	http.HandleFunc("/home", LoggingMid(ErrorMid(home)))
	http.HandleFunc("/search", LoggingMid(search))
	panic(http.ListenAndServe(":3000", nil))
	//f:= func(i int) {}

}

func LoggingMid(next http.HandlerFunc) http.HandlerFunc { // return type means exec next thing in the chain

	return func(w http.ResponseWriter, r *http.Request) {
		//do middleware specific stuff first and when it is over then go to the actual handler func to exec it
		log.Println("started")
		defer log.Println("ended")
		log.Println(r.Method)
		if r.Method != http.MethodGet {
			http.Error(w, "method allowed is GET", http.StatusBadRequest)
			return
		}

		next(w, r)

	}

}

func ErrorMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("I am handling error here")
		next(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("in home handler")
	fmt.Fprintln(w, "this is our test home page")
}

func search(w http.ResponseWriter, r *http.Request) {
	log.Println("in search")
	fmt.Fprintln(w, "this is our search page")
}
