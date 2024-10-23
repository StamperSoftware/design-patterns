package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(i int) *Pet
	SetMaxWeight(i int) *Pet
	SetWeight(i int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(i int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(i int) *Pet
	SetAgeEstimated(b bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetMinWeight(i int) *Pet {
	p.MinWeight = i
	return p
}
func (p *Pet) SetMaxWeight(i int) *Pet {
	p.MaxWeight = i
	return p
}
func (p *Pet) GetAvgWeight() int {
	return (p.MaxWeight + p.MinWeight) / 2
}
func (p *Pet) SetWeight(i int) *Pet {
	p.Weight = i
	return p
}
func (p *Pet) SetAge(i int) *Pet {
	p.Age = i
	return p
}
func (p *Pet) SetAgeEstimated(b bool) *Pet {
	p.AgeEstimated = b
	return p
}
func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}
func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}
func (p *Pet) SetDescription(s string) *Pet {
	p.Description = s
	return p
}
func (p *Pet) SetLifeSpan(i int) *Pet {
	p.LifeSpan = i
	return p
}
func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}
func (p *Pet) SetColor(s string) *Pet {
	p.Color = s
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MaxWeight < p.MinWeight {
		return nil, errors.New("max weight must be greater than min weight")
	}

	p.AvgWeight = p.GetAvgWeight()

	return p, nil
}
