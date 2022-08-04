package handlers

import (
	"context"
	"net/http"
	"service-app/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	//return web.NewRequestError(errors.New("faking the request err"), http.StatusBadRequest)
	//return errors.New("some app specific error")
	panic("some problem")
	status := struct {
		Status string //fields
	}{
		Status: "ok", //values
	}
	return web.Respond(ctx, w, status, http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//return json.NewEncoder(w).Encode(status) // convert status and write to client

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
