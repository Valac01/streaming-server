package main

import (
	"flag"
	"fmt"
)

type application struct {
	config config
}

type config struct {
	env            string
	uploadDir      string
	videoChunkSize int
}

const version = "1.0.0"

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.env, "env", "development", "server environment setting (development, production, staging)")
	flag.StringVar(&cfg.uploadDir, "upload-dir", "/uploads", "folder Name where client uploaded files will reside")
	flag.IntVar(&cfg.videoChunkSize, "video-chunk-size", 1_000_000, "size of video chunks to be sent in (bytes)")
	flag.Parse()
	app := application{
		config: *cfg,
	}

	err := app.serve()

	if err != nil {
		fmt.Println(err.Error())
	}
}
