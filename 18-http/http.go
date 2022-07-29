package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/home", homePage) // http router // handlers //
	http.HandleFunc("/search", searchSomething)
	http.ListenAndServe(":8080", nil) // this will start your sever // nil means default config // block this line
	fmt.Println("end of the main")
}

func searchSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "we are searching for something")
}

// Home http.ResponseWriter is used to write response back to the client ,
//http.Request is used to fetch any request specific details like json, body , or anything related to request data
func homePage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("welcome to our home page"))
}
