package domain

import (
	"context"
	"time"

	cityD "github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
	provD "github.com/Joskeiner/Api_e-commerce/internal/app/provice/domain"
)

// User is a struct that represents the User account.
type User struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Password    string          `json:"password"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	BirthDate   time.Time       `json:"birth_date"`
	About       string          `json:"about"`
	Job         string          `json:"job"`
	ProvinceID  string          `json:"province_id"`
	CityID      string          `json:"city_id"`
	Province    *provD.Province `json:"province"`
	City        *cityD.City     `json:"city"`
	IsAdmin     bool            `json:"is_admin"`
}

// UserRepository is an interface that provides access to the Use storage
type UserRepository interface {
	// GetByID retuns the User with the specified ID
	GetByID(ctx context.Context, id uint) (*User, error)
	// Update update the User with the specified Id
	Update(ctx context.Context, user *User) error
	// IsAdmin checks is the User with the specified ID is admin
	IsAdmin(ctx context.Context, id uint) (bool, error)
}

// UserUseCase is an interface that provide business logic for user
type UserUseCase interface {
	// GetByID returns the user with the specified id
	GetByID(ctx context.Context, id uint) (*User, error)
	// Update update the User with the specified id
	Update(ctx context.Context, user *User) error
}
