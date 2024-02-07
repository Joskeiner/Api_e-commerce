package token

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/config"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
)

// Token es an interface for token implementations
type Token interface {
	Create(userID uint, idAdmin bool) (string, error)
	Verify(token string) (*Payload, error)
}

// New Create a new token instance based on the given token type
func New(tokenCfg *config.Token) (Token, error) {
	switch tokenCfg.Type {
	case "jwt":
		return newJWT(tokenCfg.SymmetricKey, tokenCfg.Duration)
	default:
		return nil, helper.ErrUnsupportedTokenType
	}
}
