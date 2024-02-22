package usecase

import (
	"context"

	cityDom "github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
	provDom "github.com/Joskeiner/Api_e-commerce/internal/app/province/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/app/user/domain"
)

// userUseCase is a struct that implements userUseCase interface
type UserUseCase struct {
	userRepo domain.UserRepository
	cityRepo cityDom.CityRepository
	provRepo provDom.ProvinceRepository
}

func New(userRepo domain.UserRepository, cityRepo cityDom.CityRepository, provRepo provDom.ProvinceRepository) domain.UserUsecase {
	return &UserUseCase{
		userRepo,
		cityRepo,
		provRepo,
	}
}

// GetByID returns  the User with the specified ID
func (uu *UserUseCase) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := uu.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	province, err := uu.provRepo.GetByID(user.ProvinceID)
	if err != nil {
		return nil, err
	}
	city, err := uu.cityRepo.GetByID(user.ProvinceID, user.CityID)
	if err != nil {
		return nil, err
	}
	user.Province = province
	user.City = city
	return user, nil
}

// Update updates the User with the specified ID.
func (uu *UserUseCase) Update(ctx context.Context, user *domain.User) error {
	_, err := uu.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}

	return uu.userRepo.Update(ctx, user)
}

// IsAdmin checks if the User with the specified ID is an admin.
func (uu *UserUseCase) IsAdmin(ctx context.Context, id uint) (bool, error) {
	return uu.userRepo.IsAdmin(ctx, id)
}
