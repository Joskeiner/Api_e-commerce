package domain

import (
	"context"

	cityDom "github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
	provDom "github.com/Joskeiner/Api_e-commerce/internal/app/province/domain"
	userDom "github.com/Joskeiner/Api_e-commerce/internal/app/user/domain"
)

// AuthRepository is an interface that provides acces to the user storage
type AuthRepository interface {
	// create store a new User
	Create(ctx context.Context, user *userDom.User) error
	// GetbyPhoneNumber retuns the User with the specified phone number
	GetbyPhoneNumber(ctx context.Context, phoneNumber string) (*userDom.User, error)
}

// AuthUseCase is an interface that proides business logic for User
type AuthUseCase interface {
	// REgister stores a new user
	Register(ctx context.Context, user *userDom.User) error
	// login retuns User data and access token.
	Login(ctx context.Context, phoneNumber, password string) (*userDom.User, *cityDom.City, *provDom.Province, string, error)
}
