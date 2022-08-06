package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{} // Helper data structure for json responses

var pattern = regexp.MustCompile(`\D`)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) readParamID(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 0 {
		return 0, errors.New("invalid ID")
	}
	return id, nil
}

func (app *application) parseByteRange(r *http.Request) (int64, error) {
	rangeStr := r.Header.Get("range")
	if len(rangeStr) <= 0 {
		return 0, errors.New("request header required range key")
	}

	rangeStr = pattern.ReplaceAllString(rangeStr, "")
	rangeByte, err := strconv.ParseInt(rangeStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return rangeByte, nil
}

func (app *application) getVideoFile(path string) (*os.File, error) {
	return os.Open(path)
}

func (app *application) getVideoChunk(vf *os.File, r *http.Request) ([]byte, int64, error) {
	start, err := app.parseByteRange(r)
	if err != nil {
		log.Printf("Range Err %s\n", err.Error())
		return nil, 0, err
	}

	chunk := make([]byte, app.config.videoChunkSize)
	n, err := vf.ReadAt(chunk, start)
	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			return chunk[:n], start, nil
		default:
			return nil, start, err
		}
	}
	return chunk, start, nil
}

func (app *application) writeVideoChunk(w http.ResponseWriter, data *[]byte, start, vfSize int64) {
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, start+int64(len(*data)), vfSize))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(*data)))
	w.Header().Set("Content-type", "video/mp4")
	w.WriteHeader(http.StatusPartialContent)
	w.Write(*data)
}
