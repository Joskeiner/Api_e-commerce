package usecase

import "github.com/Joskeiner/Api_e-commerce/internal/app/province/domain"

// ProvinceUsecase is a struct that implements PronvinceUsecase interface
type ProvinceUseCase struct {
	repo domain.ProvinceRepository
}

// New ceates a new ProvinceUseCase instance
func New(repo domain.ProvinceRepository) domain.ProvinceRepository {
	return &ProvinceUseCase{
		repo,
	}
}

// GetAll returns all Provinces
func (pu *ProvinceUseCase) GetAll() ([]domain.Province, error) {
	return pu.repo.GetAll()
}

// GetByID returns the Province with the specified ID.
func (pu *ProvinceUseCase) GetByID(id string) (*domain.Province, error) {
	return pu.repo.GetByID(id)
}
