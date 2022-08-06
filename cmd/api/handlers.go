package main

import (
	"fmt"
	"html/template"
	"log"
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

// Render a page to stream video Will be modified to not render a html
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./uploads/html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func (app *application) videoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := app.readParamID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vf, err := app.getVideoFile(fmt.Sprintf(".%s/videos/sampleVid.mp4", app.config.uploadDir))
	if err != nil {
		log.Println(err.Error())
	}

	info, err := vf.Stat()
	if err != nil {
		log.Println(err.Error())
	}

	data, start, err := app.getVideoChunk(vf, r)
	if err != nil {
		log.Println(err.Error())
	}

	app.writeVideoChunk(w, &data, start, info.Size())
}
