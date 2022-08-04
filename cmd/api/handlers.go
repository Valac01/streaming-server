package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	serverHealth := &envelope{
		"status": "avaliable",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, serverHealth, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (app *application) videoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("video file"))
}
