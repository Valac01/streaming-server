package main

import (
	"flag"
	"fmt"
)

type application struct {
	config config
}

type config struct {
	uploadLocation string
}

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.uploadLocation, "upload-location", "/upload", "folder Name where client uploaded files will reside")
	flag.Parse()
	app := application{
		config: *cfg,
	}

	err := app.serve()

	if err != nil {
		fmt.Println(err.Error())
	}
}
