package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "1.0.0"

func main() {
	log.Printf("Starting Application of :%s Version\n", version)
	app := getApplicationConfig()
	app.startServer()

}

func getApplicationConfig() application {
	var config config
	flag.IntVar(&config.port, "port", 4000, "Server to listen on this port")
	flag.StringVar(&config.env, "env", "dev", "Environment on which the app is running (dev|prod)")
	flag.Parse()
	config.version = version
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: config,
		logger: logger,
	}
	return app
}

func (app *application) startServer() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
	err := srv.ListenAndServe()
	log.Printf("Started Server on : %d \n", app.config.port)
	if err != nil {
		log.Fatalf("Failed to Start Server on : %d \n", app.config.port)
	}
}
