package main

import (
	"fmt"
	"net/http"
)

func (app *application) serve() error {
	srv := http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	fmt.Println("Starting server at ", srv.Addr)

	return srv.ListenAndServe()
}
