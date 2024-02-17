package v1

import "github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"

// cityPrama is a struct for structuring GetbyId parameter
type cityParam struct {
	ID         string `params:"city_id"`
	ProvinceID string `params:"prov_id"`
}

// citiesParam is a struct for structuring GetAll parameter
type citiesParam struct {
	ProvinceID string `params:"prov_id"`
}

// cityResponse is a struct for structuring city response
type cityResponse struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// NewCityResponse creates a new instance of cityResponse
func NewCityResponse(city domain.City) *cityResponse {
	return &cityResponse{
		ID:         city.ID,
		ProvinceID: city.ProvinceID,
		Name:       city.Name,
	}
}
