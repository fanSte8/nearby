package main

import (
	"net/http"
)

func (app *application) handler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Data struct {
			Field string `json:"field"`
		} `json:"data"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.logger.Info("Error reading json", "error", err)
		return
	}

	app.writeJSON(w, http.StatusOK, input, http.Header{})
}
