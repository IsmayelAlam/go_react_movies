package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *application) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, values := range headers[0] {
			w.Header()[key] = values
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(output); err != nil {
		return err
	}

	return nil
}
func (app *application) readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	if err := dec.Decode(data); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app *application) errorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJson(w, statusCode, payload)
}
