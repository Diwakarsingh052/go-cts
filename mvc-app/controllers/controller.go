package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/username/reponame/model"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is our home page")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userIdString := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	if err != nil {
		// impl json responses
		http.Error(w, "please provide a valid id", http.StatusBadRequest)
		return
	}

	user, err := model.FetchUser(userId)

	if err != nil {
		http.Error(w, "please provide a valid id", http.StatusBadRequest)
		return
	}

	//jsonData, err := json.Marshal(user)
	//if err != nil {
	//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Write(jsonData) // writing the json to the client
	//w.WriteHeader(http.StatusFound) // change status code
	json.NewEncoder(w).Encode(user) // writing and converting json

}
