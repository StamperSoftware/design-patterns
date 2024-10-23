package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/test-patterns", app.TestPatterns)
	mux.Get("/api/factory-dog", app.CreateFactoryDog)
	mux.Get("/api/factory-cat", app.CreateFactoryCat)

	mux.Get("/api/abstract-factory-dog", app.CreateAbstractDog)
	mux.Get("/api/abstract-factory-cat", app.CreateAbstractCat)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	mux.Get("/api/dog-breeds", app.GetAllDogBreeds)

	return mux
}
