package main

import (
	"design-patterns/models"
	"design-patterns/pets"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
	"net/http"
	"net/url"
	"time"
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

	dogBreeds, err := app.App.Models.DogBreed.AllBreeds()

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

func (app *application) GetDogOfMonth(w http.ResponseWriter, r *http.Request) {

	var t toolbox.Tools
	breed, err := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	fmt.Println(breed)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	dogOfMonth, err := app.App.Models.Dog.GetDogOfMonthByID(1)
	fmt.Println(dogOfMonth)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	dateOfBirthLayout := "2006-01-02"
	dateOfBirth, _ := time.Parse(dateOfBirthLayout, "2020-01-01")

	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			DogName:          "Sam",
			BreedID:          breed.ID,
			Color:            "black",
			DateOfBirth:      dateOfBirth,
			SpayedOrNeutered: 0,
			Description:      "Nice Pup",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dogOfMonth.Video,
		Image: dogOfMonth.Image,
	}
	data := make(map[string]any)
	data["dog"] = dog
	app.render(w, "dog-of-the-month.page.gohtml", &templateData{Data: data})
}
