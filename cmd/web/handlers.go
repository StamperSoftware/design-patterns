package main

import (
	"design-patterns/pets"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
	"net/http"
	"net/url"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateFactoryDog(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateFactoryCat(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateAbstractDog(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.CreateAbstractFactoryPet("dog")

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateAbstractCat(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.CreateAbstractFactoryPet("cat")

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	dogBreeds, err := app.App.Models.DogBreed.All()

	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateBuilderDog(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed").
		SetWeight(15).
		SetDescription("Cool pup").
		SetColor("Brown").
		SetAge(5).
		SetAgeEstimated(false).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateBuilderCat(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("calico").
		SetWeight(5).
		SetDescription("Cool cool cat").
		SetColor("orange").
		SetAge(9).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetRemoteCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	catBreeds, err := app.App.CatService.GetAllBreeds()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) CreateAbstractAnimal(w http.ResponseWriter, r *http.Request) {

	var t toolbox.Tools
	species := chi.URLParam(r, "species")
	breed := chi.URLParam(r, "breed")
	breed, _ = url.QueryUnescape(breed)

	animal, err := pets.CreateAbstractFactoryPetWithBreed(species, breed)
	
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, animal)
}
