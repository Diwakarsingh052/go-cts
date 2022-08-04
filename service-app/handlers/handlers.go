package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-app/middleware"
	"service-app/web"
)

func API(log *log.Logger) http.Handler {
	//app := mux.NewRouter()
	app := web.App{Router: mux.NewRouter()}
	m := middleware.Mid{Log: log}

	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(check))))

	return app
}
