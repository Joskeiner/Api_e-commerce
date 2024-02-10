package repository

import (
	"encoding/json"
	"os"

	"github.com/Joskeiner/Api_e-commerce/internal/app/province/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
)

// ProvinceRepository is a struct that implements ProvinceRepository interface
type ProvinceRepositoryJSON struct {
	Province map[string]domain.Province
}

func New() domain.ProvinceRepository {
	return &ProvinceRepositoryJSON{
		Province: make(map[string]domain.Province),
	}
}

// basePath is a constant that represents the base path of the API endpoint
const basePath = "../../../pkg/database/infoAR/"

// GetAll is a method that returns all province
func (pr *ProvinceRepositoryJSON) GetAll() ([]domain.Province, error) {
	url := basePath + "provinceAR.json"

	content, err := os.ReadFile(url)
	if err != nil {
		return nil, helper.ErrDataNotFound
	}
	provinces, err := pr.ParseJSON(content)
	if err != nil {
		return nil, err
	}
	return *provinces, nil
}

// GetByID is a method that returns the province with the specified ID.
func (pr *ProvinceRepositoryJSON) GetByID(id string) (*domain.Province, error) {
	prov, ok := pr.Province[id]
	if !ok {

		url := basePath + "provinceAR.json"
		content, err := os.ReadFile(url)
		if err != nil {
			return nil, helper.ErrDataNotFound
		}
		provinces, err := pr.ParseJSON(content)
		if err != nil {
			return nil, err
		}
		for _, prov := range *provinces {
			pr.Province[prov.ID] = prov
		}
		province, ok := pr.Province[id]
		if !ok {
			return nil, helper.ErrDataNotFound
		}

		return &province, nil
	}
	return &prov, nil
}

func (pr *ProvinceRepositoryJSON) ParseJSON(provJson []byte) (*[]domain.Province, error) {
	var province []domain.Province
	err := json.Unmarshal(provJson, &province)
	if err != nil {
		return nil, err
	}

	return &province, nil
}
