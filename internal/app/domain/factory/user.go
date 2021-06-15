package factory

import (
	"github.com/airdb/uic/internal/app/domain"
)

// Order is the factory of domain.User
type User struct{}

// Generate generates domain.User from primitives
func (of User) Generate(
	id uint64,
	username string,
) domain.User {
	return domain.User{
		ID:       id,
		Username: username,
	}
}

type OauthConfig struct{}

func (of OauthConfig) GenerateOauthConfig(id uint, clientID, redirectURL, state string) domain.OauthConfig {
	return domain.OauthConfig{
		ID:          id,
		ClientID:    clientID,
		RedirectURL: redirectURL,
		State:       state,
	}
}
