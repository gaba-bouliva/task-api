package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"task-api/internal/data"
	"time"
)

const version = "1.0.0"

type config struct {
	port 		int
	env     string
}

type application struct {
	models 			data.Models
	config			config
	logger 			*log.Logger
	version 		string
}



func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8081, "API Server port")
	flag.StringVar(&cfg.env, "env", "development", "[ development | staging | production ]")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)

	app := application{
		models: data.NewModels(nil), // TODO replace the nil with a real db connection i.e *sql.DB
		config: cfg,
		logger: logger,
		version: version,
	}

	if err := app.run(); err != nil {
		logger.Fatal(err)
	}
}

func (app *application) run() error{
	srv := http.Server{
		Handler: app.routes(),
		Addr: fmt.Sprintf(":%d",app.config.port),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Starting %s server on port %d", app.config.env, app.config.port)
	return srv.ListenAndServe() 
}