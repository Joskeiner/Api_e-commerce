package domain

import "context"

type Shop struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	UserID         uint   `json:"user_id"`
}

// ShopRepository is an interface that provides access to the Shop Storage
type ShopRepository interface {
	GetAll(ctx context.Context, page, limit int, name string) ([]Shop, error)

	GetUserShop(ctx context.Context, userID uint) (*Shop, error)

	GetByID(ctx context.Context, id uint) (*Shop, error)

	Update(ctx context.Context, shop *Shop) error
}

// ShopUseCase is an interface that provides business logic for shop
type ShopUseCase interface {
	GetAll(ctx context.Context, page, limit int, name string) ([]Shop, error)

	GetUserShop(ctx context.Context, userID uint) (*Shop, error)

	GetByID(ctx context.Context, id uint) (*Shop, error)

	Update(ctx context.Context, shop *Shop) error
}
