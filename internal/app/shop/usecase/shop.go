package usecase

import (
	"context"

	"github.com/Joskeiner/Api_e-commerce/internal/app/shop/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
)

// ShopUsecase is a struct that implements ShopUsecase interface.
type ShopUsecase struct {
	repo domain.ShopRepository
}

func New(repo domain.ShopRepository) domain.ShopUseCase {
	return &ShopUsecase{
		repo,
	}
}

// GetAll() return all shop
func (su *ShopUsecase) GetAll(ctx context.Context, page, limit int, name string) ([]domain.Shop, error) {
	return su.repo.GetAll(ctx, page, limit, name)
}

// GetUserShop() returns the shop with the specified user id
func (su *ShopUsecase) GetUserShop(ctx context.Context, userID uint) (*domain.Shop, error) {
	return su.repo.GetUserShop(ctx, userID)
}

// GetByID returns the shop the specified id
func (su *ShopUsecase) GetByID(ctx context.Context, id uint) (*domain.Shop, error) {
	return su.repo.GetByID(ctx, id)
}

// Update() update the Shop with the specified id
func (su *ShopUsecase) Update(ctx context.Context, userID, id uint) error {
	shop, err := su.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if shop.UserID != userID {
		return helper.ErrForbidden
	}
	return su.repo.Update(ctx, shop)
}
