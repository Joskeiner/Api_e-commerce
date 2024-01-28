package v1

import (
	"github.com/Joskeiner/Api_e-commerce/internal/app/shop/domain"
)

// shopParam is a struct for validating shop path parameter
type shopParam struct {
	ID uint `params:"id" validate:"required"`
}

// shopQuery is a struct for validating shop query parameter

type shopQuery struct {
	Page  int    `form:"page" validate:"required,min=1"`
	Limit int    `form:"limit" validate:"required,min=5"`
	Name  string `form:"name" validate:"omitempty"`
}

// updateShopRequest is a struct for validating update shop request body
type updateShopResquest struct {
	Name           string `form:"name" validate:"omitempty"`
	ProfilePicture string `form:"profile_picture" validate:"omitempty"`
}

// shopResponse is a struct for structuring shop response
type shopResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
}

// shopResponsePaginated is a struct for structuring paginated shop response
type shopResponsePaginated struct {
	Shops []shopResponse `json:"shops"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

func NewShopRespose(shop domain.Shop) *shopResponse {
	return &shopResponse{
		ID:             shop.ID,
		Name:           shop.Name,
		ProfilePicture: shop.ProfilePicture,
	}
}
