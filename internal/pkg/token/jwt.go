package token

import (
	"errors"
	"time"

	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/golang-jwt/jwt/v4"
)

const minScretKeySize = 32

type Jwt struct {
	secretKey string
	duration  time.Duration
}

// newJWT creates a new JWT token generator
func newJWT(secretKey string, duration time.Duration) (*Jwt, error) {
	if len(secretKey) < minScretKeySize {
		return nil, helper.ErrInvalidToken
	}
	return &Jwt{
		secretKey,
		duration,
	}, nil
}

// Create creates a new token for a specific user id, role , and duration
func (j *Jwt) Create(userID uint, isAdmin bool) (string, error) {
	payload, err := NewPayload(userID, isAdmin, j.duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	token, err := jwtToken.SignedString([]byte(j.secretKey))

	return token, nil
}

// verify checks if the token is valid or not
func (j *Jwt) verify(token string) (*Payload, error) {
	KeyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, helper.ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, KeyFunc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr.Inner, helper.ErrExpiredToken) {
			return nil, helper.ErrExpiredToken
		}
		return nil, helper.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, helper.ErrInvalidToken
	}
	return payload, nil
}
