package product_handlers

import (
	"investidea.tech.test/internal/repository"
	"logur.dev/logur"
)

type productHandler struct {
	logger logur.LoggerFacade
	repo   repository.Registry
}

func New(logger logur.LoggerFacade, repo repository.Registry) productHandler {
	return productHandler{
		logger: logger,
		repo:   repo,
	}
}
