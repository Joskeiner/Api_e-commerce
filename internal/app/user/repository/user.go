package repository

import (
	"context"

	"github.com/Joskeiner/Api_e-commerce/internal/app/user/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database/dao"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
)

// UserRepository is a struct that implements UserRepository interface
type UserRepository struct {
	conn database.DB
}

// New creates a new UserRepository instance
func New(conn database.DB) domain.UserRepository {
	return &UserRepository{
		conn,
	}
}

// GetByID returns the User with the specidied Id
func (ur *UserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var dao dao.User

	result := ur.conn.DB().First(&dao, id).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, helper.ErrDataNotFound
		}

		return nil, result.Error
	}

	user := ur.toDomain(&dao)

	return user, nil
}

// Update updates the User with the specified ID.
func (ur *UserRepository) Update(ctx context.Context, user *domain.User) error {
	dao := ur.toDAO(user)

	result := ur.conn.DB().Model(dao).Updates(&dao).Where("id = ?", dao.ID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return helper.ErrDataNotFound
		}

		return result.Error
	}

	return nil
}

func (ur *UserRepository) IsAdmin(ctx context.Context, id uint) (res bool, err error) {
	var dao dao.User
	result := ur.conn.DB().Where("id = ?", id).First(&dao).WithContext(ctx)

	if result.Error != nil {
		return false, helper.ErrDataNotFound
	}
	return dao.IsAdmin, nil
}

// toDamain converts a DAO category to a category
func (ur *UserRepository) toDomain(user *dao.User) *domain.User {
	return &domain.User{
		ID:          user.ID,
		Name:        user.Name,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		BirthDate:   user.BirthDate,
		About:       user.About,
		Job:         user.Job,
		ProvinceID:  user.ProvinceID,
		CityID:      user.CityID,
	}
}

// toDao converts a category to DAO category
func (ur *UserRepository) toDAO(user *domain.User) *dao.User {
	return &dao.User{
		Model: dao.Model{
			ID: user.ID,
		},
		Name:        user.Name,
		Password:    user.Password,
		Email:       user.Email,
		BirthDate:   user.BirthDate,
		About:       user.About,
		Job:         user.Job,
		CityID:      user.CityID,
		ProvinceID:  user.ProvinceID,
		PhoneNumber: user.PhoneNumber,
	}
}
