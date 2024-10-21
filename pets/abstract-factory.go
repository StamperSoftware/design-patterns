package pets

import (
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
