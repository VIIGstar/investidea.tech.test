package session_handlers

import (
	"investidea.tech.test/internal/repository"
	"logur.dev/logur"
)

type sessionHandler struct {
	logger logur.LoggerFacade
	repo   repository.Registry
}

func New(logger logur.LoggerFacade, repo repository.Registry) sessionHandler {
	return sessionHandler{
		logger: logger,
		repo:   repo,
	}
}
