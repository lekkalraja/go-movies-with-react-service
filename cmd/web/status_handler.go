package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     app.config.version,
	}
	data, err := json.Marshal(status)
	if err != nil {
		app.logger.Printf("Failed to Parse Status Json : %v", status)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
