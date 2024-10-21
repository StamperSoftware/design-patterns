package pets

import "design-patterns/models"

func NewPet(species string) *models.Pet {

	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "not set",
		LifeSpan:    0,
	}

	return &pet
}
