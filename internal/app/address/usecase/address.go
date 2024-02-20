package usecase

import (
	"context"

	"github.com/Joskeiner/Api_e-commerce/internal/app/address/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
)

// AddressUsecase is a struct that implements AddressUsecase interface.
type AddressUsecase struct {
	addrRepo domain.AddressRepository
}

// New creates a new CategoryUsecase instance.
func New(addrRepo domain.AddressRepository) domain.AddressUseCase {
	return &AddressUsecase{
		addrRepo,
	}
}

// Create stores a new Address.
func (au *AddressUsecase) Create(ctx context.Context, addr *domain.Address) error {
	return au.addrRepo.Create(ctx, addr)
}

// GetAll returns all Addresses.
func (au *AddressUsecase) GetAll(ctx context.Context, userID uint, title string) ([]domain.Address, error) {
	return au.addrRepo.GetAll(ctx, userID, title)
}

// GetByID retuns the address with the specified ID.
func (au *AddressUsecase) GetByID(ctx context.Context, userID, id uint) (*domain.Address, error) {
	addr, err := au.addrRepo.GetByID(ctx, userID, id)
	if err != nil {
		return nil, err
	}

	if addr.UserID != userID {
		return nil, helper.ErrForbidden
	}
	return addr, nil
}

// Update updates the Address with the specified ID.
func (au *AddressUsecase) Update(ctx context.Context, addr *domain.Address) error {
	currentAddr, err := au.GetByID(ctx, addr.UserID, addr.ID)
	if err != nil {
		return err
	}

	if currentAddr.UserID != addr.UserID {
		return helper.ErrForbidden
	}

	return au.addrRepo.Update(ctx, addr)
}

// Delete removes the Address with the specified ID.
func (au *AddressUsecase) Delete(ctx context.Context, userID, id uint) error {
	addr, err := au.GetByID(ctx, userID, id)
	if err != nil {
		return err
	}

	if addr.UserID != userID {
		return helper.ErrForbidden
	}

	return au.addrRepo.Delete(ctx, userID, id)
}
