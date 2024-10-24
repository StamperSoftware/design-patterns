package pets

import (
	"design-patterns/configuration"
	"design-patterns/models"
	"errors"
	"fmt"
)

type AnimalInterface interface {
	Show() string
}

type FactoryDog struct {
	Pet *models.Dog
}

func (fd *FactoryDog) Show() string {
	return fmt.Sprintf("This animal is a %s", fd.Pet.Breed.Breed)
}

type FactoryCat struct {
	Pet *models.Cat
}

func (fc *FactoryCat) Show() string {
	return fmt.Sprintf("This cat is a %s", fc.Pet.Breed.Breed)
}

type FactoryPetInterface interface {
	newPet() AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

type AbstractFactoryDog struct{}

func (fd *AbstractFactoryDog) newPet() AnimalInterface {
	return &FactoryDog{
		Pet: &models.Dog{},
	}
}

type AbstractFactoryCat struct{}

func (fc *AbstractFactoryCat) newPet() AnimalInterface {
	return &FactoryCat{
		Pet: &models.Cat{},
	}
}

func (fc *AbstractFactoryDog) newPetWithBreed(b string) AnimalInterface {

	app := configuration.GetInstance()
	breed, err := app.Models.DogBreed.GetBreedByName(b)

	if err != nil {
		return nil
	}

	return &FactoryDog{
		Pet: &models.Dog{
			Breed: *breed,
		},
	}
}
func (fc *AbstractFactoryCat) newPetWithBreed(b string) AnimalInterface {

	app := configuration.GetInstance()
	breed, err := app.CatService.Remote.GetBreedByName(b)

	if err != nil {
		return nil
	}

	return &FactoryCat{
		Pet: &models.Cat{
			Breed: *breed,
		},
	}
}

func CreateAbstractFactoryPet(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var factoryDog AbstractFactoryDog
		dog := factoryDog.newPet()
		return dog, nil

	case "cat":
		var factoryCat AbstractFactoryCat
		cat := factoryCat.newPet()
		return cat, nil

	default:
		return nil, errors.New("invalid species")
	}
}

func CreateAbstractFactoryPetWithBreed(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory AbstractFactoryDog
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		var catFactory AbstractFactoryCat
		cat := catFactory.newPetWithBreed(breed)
		return cat, nil
	default:
		return nil, errors.New("invalid species")
	}
}
