package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

// use this for later
type config struct {
	port int
	env  string
}

/*
	bellow, the line is read as,
	the Status property is a string and when being
	displayed as json be displayed as "status"
*/
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	var cfg config

	// command line flags are a common way to specify options for command-line
	// programs for example in wc -l the -l is a command line flag
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse() // you parse the flags when they're created.

	fmt.Println("Running") // just to test this is working

	// basic web sever creation

	// These is how we create routes and handler functions
	// for when the server is hit at that route
	// all handlers need a response function
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// let's return some json

		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		// the json will be stored in js
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		// when writing the response, we have to specify the type,
		// the http header and then the content.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)

	if err != nil {
		log.Println(err)
	}

}
