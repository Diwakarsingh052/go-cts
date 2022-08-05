package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/auth"
	"service-app/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "user : ", log.LstdFlags)
	err := startApp(l)
	if err != nil {
		panic(err)
	}
}

func startApp(log *log.Logger) error {

	// =========================================================================
	// Initialize authentication support
	log.Println("main : Started : Initializing authentication support")

	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	a, err := auth.NewAuth(privateKey)

	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		Handler:      handlers.API(log, a),
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	//this case exec in case of problem while starting the app
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)

		//receiving over shutdown chan // if this case is true then someone pressed ctrl+c
	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		// Asking listener to shutdown and shed load.
		err := api.Shutdown(ctx) // first trying to cleanly shutdown
		if err != nil {
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}

	return nil

}
