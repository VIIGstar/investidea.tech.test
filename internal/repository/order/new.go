package order

import (
	"context"
	"investidea.tech.test/internal/entities"
	query_params "investidea.tech.test/internal/query-params"
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
	Create(order *entities.Order) error
	// Get retrieves a impl based on search criteria
	Get(ctx context.Context, req query_params.OrderParams, lock bool) (*entities.Order, error)
	// Find retrieves a impl based on search criteria
	Find(ctx context.Context, req query_params.OrderParams, lock bool) ([]entities.Order, error)
}

type impl struct {
	logger logur.LoggerFacade
	db     *database.DB
}

func (i impl) Get(ctx context.Context, req query_params.OrderParams, lock bool) (*entities.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (i impl) Find(ctx context.Context, req query_params.OrderParams, lock bool) ([]entities.Order, error) {
	//TODO implement me
	panic("implement me")
}
