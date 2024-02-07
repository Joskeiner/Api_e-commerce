package token

import (
	"time"

	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/google/uuid"
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uint      `json:"user_id"`
	IsAdmin   bool      `json:"is_admin"`
	IsUsedAt  time.Time `json:"isused_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific user id , role , and duration
func NewPayload(userID uint, isAdmin bool, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IsAdmin:   isAdmin,
		IsUsedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// valid checks if the token payload is valid or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return helper.ErrExpiredToken
	}
	return nil
}
