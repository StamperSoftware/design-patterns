package main

import (
	"design-patterns/adapters"
	"design-patterns/configuration"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
}

type appConfig struct {
	useTemplateCache bool
	dsn              string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	err := godotenv.Load()
	app.config.dsn = os.Getenv("DSN")

	if err != nil {
		log.Fatal("could not load env file")
	}

	flag.BoolVar(&app.config.useTemplateCache, "template-cache", false, "Use template cache")
	flag.Parse()

	//get db
	db, err := initMySQLDB(app.config.dsn)

	if err != nil {
		log.Panic(err)
	}

	jsonBackend := &adapters.JSONBackend{}
	jsonAdapter := &adapters.RemoteService{Remote: jsonBackend}

	//xmlBackend := &adapters.XMLBackend{}
	//xmlAdapter := &adapters.RemoteService{Remote: xmlBackend}

	app.App = configuration.New(db, jsonAdapter)
	defer db.Close()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting on port", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
