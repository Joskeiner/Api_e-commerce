package domain

// City is a struct that representes the User's addres city
type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// CityRepository is an interface that provides acces to the city storage
type CityRepository interface {
	GetAll(provinceID string) ([]City, error)
	GetByID(provinceID, id string) (*City, error)
}

// CityUsecase is an interface that provides business logic for city
type CityUseCase interface {
	// GetAll() returns all cities
	GetAll(privinceID string) ([]City, error)
	GetByID(provinceID, cityId string) (*City, error)
}
