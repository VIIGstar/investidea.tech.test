package repository

import (
	"investidea.tech.test/internal/repository/user"
	"investidea.tech.test/internal/services/database"
	"logur.dev/logur"
)

type DatabaseRepo interface {
	User() user.Repo
}

func NewDBImpl(logger logur.LoggerFacade, db *database.DB) dbImpl {
	return dbImpl{
		user: user.New(logger, db),
	}
}

type dbImpl struct {
	user user.Repo
}

func (i dbImpl) User() user.Repo {
	return i.user
}
