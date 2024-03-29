package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/double", doubleHandler)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this is the home"))
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("v") // trying to fetch value from the user
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return // don't forget the return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)

	// the return values must be tested // here we have 3

}
