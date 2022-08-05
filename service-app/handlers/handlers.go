package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/middleware"
	"service-app/web"
)

func API(log *log.Logger, A *auth.Auth) http.Handler {
	//app := mux.NewRouter()
	app := web.App{Router: mux.NewRouter()}
	m := middleware.Mid{
		Log: log,
		A:   A,
	}

	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(check)))))

	return app
}
