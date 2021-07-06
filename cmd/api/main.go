package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

/*
	Downloaded a 3rd party router from:
	github.com/julienschmidt/httprouter
	had to run
	go get -u github.com/julienschmidt/httprouter
	on terminal

	it is well suited for the purposes of this app
*/

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

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// command line flags are a common way to specify options for command-line
	// programs for example in wc -l the -l is a command line flag
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse() // you parse the flags when they're created.

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// this will be used as a receiver for other applications
	app := &application{
		config: cfg,
		logger: logger,
	}

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port:", cfg.port)

	err := serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
