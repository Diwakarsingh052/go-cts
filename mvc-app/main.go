package main

import (
	"github.com/username/reponame/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/users", controllers.GetUser)

	panic(http.ListenAndServe(":8080", nil))
}
