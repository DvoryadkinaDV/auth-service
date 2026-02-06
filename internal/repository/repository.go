package repository

import (
	"github.com/dvoryadkinadv/auth-service/pkg/database"
)

// Repositories holds all repository interfaces
type Repositories struct {
	User          UserRepository
	Token         TokenRepository
	OAuthProvider OAuthProviderRepository
}

// NewRepositories creates all repositories
func NewRepositories(db *database.Postgres) *Repositories {
	return &Repositories{
		User:          NewUserRepository(db),
		Token:         NewTokenRepository(db),
		OAuthProvider: NewOAuthProviderRepository(db),
	}
}
