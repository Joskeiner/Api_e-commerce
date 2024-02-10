package domain

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProvinceRepository is an interface that provides access to the Province storage
type ProvinceRepository interface {
	// GetAll returns all Provinces
	GetAll() ([]Province, error)

	// GetByID returns the province with the specified id
	GetByID(id string) (*Province, error)
}

// ProvinceUseCase is an interface that provides business logic for Province
type ProvinceUseCase interface {
	// GetAll returns all Province
	GetAll() ([]Province, error)

	// GetByID returns the Province with the specified id
	GetByID(id string) (*Province, error)
}
