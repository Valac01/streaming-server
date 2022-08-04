package main

import (
	"flag"
	"fmt"
)

type application struct {
	config config
}

type config struct {
	env       string
	uploadDir string
}

const version = "1.0.0"

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.env, "env", "development", "server environment setting (development, production, staging)")
	flag.StringVar(&cfg.uploadDir, "upload-Dir", "/upload", "folder Name where client uploaded files will reside")
	flag.Parse()
	app := application{
		config: *cfg,
	}

	err := app.serve()

	if err != nil {
		fmt.Println(err.Error())
	}
}
