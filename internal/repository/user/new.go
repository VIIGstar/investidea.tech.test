package user

import (
	"context"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/internal/services/database"
	"logur.dev/logur"
)

// New creates new impl impl and returns as User interface
func New(logger logur.LoggerFacade, db *database.DB) Repo {
	return &impl{
		logger: logger,
		db:     db,
	}
}

// Repo represents methods that User repository must implement
type Repo interface {
	// Create inserts new record in User table
	Create(u *entities.User) error
	Login(ctx context.Context, email, password string) (*entities.User, error)
}

type impl struct {
	logger logur.LoggerFacade
	db     *database.DB
}
