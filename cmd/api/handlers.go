package main

import "net/http"

func (app *application) HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}
