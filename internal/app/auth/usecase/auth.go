package usercase

import (
	"context"

	"github.com/Joskeiner/Api_e-commerce/internal/app/auth/domain"
	cityDom "github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
	provDom "github.com/Joskeiner/Api_e-commerce/internal/app/province/domain"
	userDom "github.com/Joskeiner/Api_e-commerce/internal/app/user/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
)

// AuthUsecase is a struct that implements AuthUsecase interface.
type AuthUseCase struct {
	authRepo   domain.AuthRepository
	userRepo   userDom.UserRepository
	cityRepo   cityDom.CityRepository
	provRepo   provDom.ProvinceRepository
	tokenMaker token.Token
}

func New(authRepo domain.AuthRepository, userRepo userDom.UserRepository, cityRepo cityDom.CityRepository, provRepo provDom.ProvinceRepository, tokenMaker token.Token) domain.AuthUseCase {
	return &AuthUseCase{
		authRepo,
		userRepo,
		cityRepo,
		provRepo,
		tokenMaker,
	}
}

// Register stores a new User
func (au *AuthUseCase) Register(ctx context.Context, user *userDom.User) error {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return au.authRepo.Create(ctx, user)
}

// Login retuns User Data and acces token
func (au *AuthUseCase) Login(ctx context.Context, phoneNumber, password string) (*userDom.User, *cityDom.City, *provDom.Province, string, error) {
	user, err := au.authRepo.GetbyPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, nil, nil, "", err
	}
	if err := helper.ComparePassword(password, user.Password); err != nil {
		return nil, nil, nil, "", err
	}
	city, err := au.cityRepo.GetByID(user.Province.ID, user.CityID)
	if err != nil {
		return nil, nil, nil, "", err
	}

	province, err := au.provRepo.GetByID(user.ProvinceID)
	if err != nil {
		return nil, nil, nil, "", err
	}

	token, err := au.tokenMaker.Create(user.ID, user.IsAdmin)
	if err != nil {
		return nil, nil, nil, "", err
	}
	return user, city, province, token, nil
}
