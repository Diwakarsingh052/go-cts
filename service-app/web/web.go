package web

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type ctxKey int

const KeyValue ctxKey = 1

type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

type App struct {
	*mux.Router
}

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) HandleFunc(method string, path string, handler HandlerFunc) {

	h := func(w http.ResponseWriter, r *http.Request) {

		v := Values{
			TraceID: uuid.New().String(),
			Now:     time.Now(),
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, KeyValue, &v)

		err := handler(ctx, w, r)

		if err != nil {
			log.Println(err)
			return
		}
	}

	a.Router.HandleFunc(path, h).Methods(method)

}
