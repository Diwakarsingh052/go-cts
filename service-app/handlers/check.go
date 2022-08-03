package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	status := struct {
		Status string //fields
	}{
		Status: "ok", //values
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(status) // convert status and write to client

}

//v1
//func check( w http.ResponseWriter, r *http.Request)  {
//
//	status := struct {
//		Status string //fields
//	}{
//		Status: "ok", //values
//	}
//
//	json.NewEncoder(w).Encode(status) // convert status and write to client
//
//}
