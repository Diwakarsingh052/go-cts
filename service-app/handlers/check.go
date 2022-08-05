package handlers

import (
	"context"
	"net/http"
	"service-app/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	//return web.NewRequestError(errors.New("faking the request err"), http.StatusBadRequest)
	//return errors.New("some app specific error")

	status := struct {
		Status string //fields
	}{
		Status: "ok", //values
	}
	//claims, ok := ctx.Value(auth.Key).(auth.Claims)
	//if !ok {
	//	return errors.New("not found")
	//}

	return web.Respond(ctx, w, status, http.StatusOK)

	//ctx := r.Context()
	//model.SetUser(ctx,userStruct)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(statusCode)
	// json.NewEncoder(w).Encode(status) // convert status and write to client

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
