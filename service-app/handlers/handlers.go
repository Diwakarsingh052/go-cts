package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/data/user"
	"service-app/middleware"
	"service-app/web"
)

func API(log *log.Logger, A *auth.Auth, db *user.DbService) http.Handler {
	//app := mux.NewRouter()
	app := web.App{Router: mux.NewRouter()}
	m := middleware.Mid{
		Log: log,
		A:   A,
	}
	uh := userHandlers{
		DbService: db,
		Auth:      A,
	}
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.HasRole(check, auth.RoleAdmin))))))
	app.HandleFunc(http.MethodPost, "/create", m.Logger(m.Error(m.Panic(uh.SignUp))))
	app.HandleFunc(http.MethodPost, "/login", m.Logger(m.Error(m.Panic(uh.Login))))
	app.HandleFunc(http.MethodGet, "/add", m.Logger(m.Error(m.Panic(m.Authenticate(uh.AddInventory)))))
	app.HandleFunc(http.MethodGet, "/view", m.Logger(m.Error(m.Panic(m.Authenticate(uh.ViewInventory)))))
	return app
}
