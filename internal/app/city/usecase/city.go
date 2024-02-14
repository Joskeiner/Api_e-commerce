package usecase

import (
	"github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
)

// CityUseCase is a struct that implemenst CityUseCase interface
type CityUseCase struct {
	repo domain.CityRepository
}

// New creates a new CityUseCase instance
func New(repo domain.CityRepository) domain.CityUseCase {
	return &CityUseCase{
		repo,
	}
}

// GetAll retuns the City with the specified Id.
func (cu *CityUseCase) GetAll(provinceID string) ([]domain.City, error) {
	return cu.repo.GetAll(provinceID)
}

// GetByID returns the city with the specified ID.
func (cu *CityUseCase) GetByID(provinceID, CityId string) (*domain.City, error) {
	return cu.repo.GetByID(provinceID, CityId)
}
