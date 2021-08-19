package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	log.Printf("Starting Application of :%s Version\n", version)
	config := parseArgs()
	startServer(config)

}

func parseArgs() config {
	var config config
	flag.IntVar(&config.port, "port", 4000, "Server to listen on this port")
	flag.StringVar(&config.env, "env", "dev", "Environment on which the app is running (dev|prod)")
	flag.Parse()
	return config
}

func startServer(config config) {
	http.HandleFunc("/status", statusHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil)
	log.Printf("Started Server on : %d \n", config.port)
	if err != nil {
		log.Fatalf("Failed to Start Server on : %d \n", config.port)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := AppStatus{
		Status:      "Available",
		Environment: "Dev",
		Version:     version,
	}
	data, err := json.Marshal(status)
	if err != nil {
		log.Printf("Failed to Parse Status Json : %v", status)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
