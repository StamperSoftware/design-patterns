package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreeds(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreeds)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code. Got %d instead of 200", rr.Code)
	}

}

func TestApplication_GetRemoteCatBreeds(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetRemoteCatBreeds)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code. Got %d instead of 200", rr.Code)
	}

}
