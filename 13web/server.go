package main

import (
	"net/http"
	"time"
)

func (app *application) Serve() error {

	srv := http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      app.routes(),
	}

	return srv.ListenAndServe()
}
