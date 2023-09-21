package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cvpchatalert/pkg/jsonlog"
)

//Create a main struct for the configuration

type Application struct {
	Port   int
	logger *jsonlog.Logger
}

// Create a envelope to write the json that takes any structure

type Envelope map[string]any

// Write JSON method will write the json to the http out body

func (app *Application) writeJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
} 

// Starts the API

func (app *Application) StartApi() {
	// Logging
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// Struct to hold the http server info
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// Start the HTTP server.
	logger.Printf("starting rest server on http://localhost:%v", app.Port)
	HttpStart := srv.ListenAndServe()
	logger.Fatal(HttpStart)
}
