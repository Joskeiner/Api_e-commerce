package repository

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
)

// CityRepositoryJSON is a struct that implements CityRepositoryJSON interface
type CityRepositoryJSON struct {
	Cities map[string]domain.City
}

// base path is a constant that represents the base path of the API endpoint
const basePath = "internal/pkg/database/infoAR/cityAR/"

// New is function that returns new CityRepositoryJSON instance
func New() domain.CityRepository {
	return &CityRepositoryJSON{
		Cities: make(map[string]domain.City),
	}
}

// GetAll is a method that returns all Cities
func (cr *CityRepositoryJSON) GetAll(provinceID string) ([]domain.City, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, helper.ErrDataNotFound
	}
	url := filepath.Join(path, basePath, provinceID+".json")
	content, err := os.ReadFile(url)
	if err != nil {
		return nil, helper.ErrDataNotFound
	}
	cities, err := cr.ParseJSON(content)
	if err != nil {
		return nil, err
	}
	return *cities, nil
}

// GetByID is a method that returns the City with the specified ID.
func (cr *CityRepositoryJSON) GetByID(provinceID, id string) (*domain.City, error) {
	city, ok := cr.Cities[id]
	if !ok {
		cities, err := cr.GetAll(provinceID)
		if err != nil {
			return nil, helper.ErrDataNotFound
		}
		for _, c := range cities {
			cr.Cities[c.ID] = c
		}
		cityR, ok := cr.Cities[id]
		if !ok {
			return nil, helper.ErrDataNotFound
		}
		return &cityR, nil
	}
	return &city, nil
}

func (cr *CityRepositoryJSON) ParseJSON(citiesJSON []byte) (*[]domain.City, error) {
	var cities []domain.City
	err := json.Unmarshal(citiesJSON, &cities)
	if err != nil {
		return nil, err
	}

	return &cities, nil
}
