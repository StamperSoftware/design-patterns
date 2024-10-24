package adapters

import (
	"design-patterns/models"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
	GetBreedByName(b string) (*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsInterface
}

func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}
func (rs *RemoteService) GetBreedByName(b string) (*models.CatBreed, error) {
	return rs.Remote.GetBreedByName(b)
}

type JSONBackend struct{}

func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed
	err = json.Unmarshal(body, &breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil
}

func (jb *JSONBackend) GetBreedByName(b string) (*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var breed models.CatBreed
	err = json.Unmarshal(body, &breed)
	if err != nil {
		return nil, err
	}

	return &breed, nil

}

type XMLBackend struct{}

func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds

	err = xml.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds.Breeds, nil

}

func (xb *XMLBackend) GetBreedByName(b string) (*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/xml")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var breed models.CatBreed
	err = xml.Unmarshal(body, &breed)
	if err != nil {
		return nil, err
	}

	return &breed, nil

}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{
			ID:      1,
			Breed:   "Tomcat",
			Details: "cool cat",
		},
	}

	return breeds, nil
}

func (tb *TestBackend) GetBreedByName(b string) (*models.CatBreed, error) {
	return nil, nil
}