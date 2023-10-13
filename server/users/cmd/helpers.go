package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	app.logger.Info("Sending response", "data", data)

	w.Header().Set("Content-Type", "application/json")

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.WriteHeader(status)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := json.NewDecoder(r.Body).Decode(dst)

	if err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		default:
			return err
		}
	}

	app.logger.Info("Received json", "data", dst)

	return nil
}
