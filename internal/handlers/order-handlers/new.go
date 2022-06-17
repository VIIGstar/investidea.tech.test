package order_handlers

import (
	"investidea.tech.test/internal/repository"
	"logur.dev/logur"
)

type orderHandler struct {
	logger logur.LoggerFacade
	repo   repository.Registry
}

func New(logger logur.LoggerFacade, repo repository.Registry) orderHandler {
	return orderHandler{
		logger: logger,
		repo:   repo,
	}
}
