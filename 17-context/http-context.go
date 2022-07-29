package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", randomData)
	http.ListenAndServe(":8080", nil)
}

func randomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("started")
	defer log.Println("ended")
	select {
	case <-time.After(1 * time.Millisecond):
		fmt.Fprintln(w, "random data")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx.Err())
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
