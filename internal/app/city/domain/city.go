package domain

// City is a struct that representes the User's addres city
type City struct {
	ID         string `json:"id"`
	provinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// CityRepository is an interface that provides acces to the city storage

// CityUsecase is an interface that provides business logic for city
