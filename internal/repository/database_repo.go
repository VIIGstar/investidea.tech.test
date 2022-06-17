package repository

import (
	"investidea.tech.test/internal/repository/order"
	"investidea.tech.test/internal/repository/product"
	"investidea.tech.test/internal/repository/user"
	"investidea.tech.test/internal/services/database"
	"logur.dev/logur"
)

type DatabaseRepo interface {
	User() user.Repo
	Product() product.Repo
	Order() order.Repo
}

func NewDBImpl(logger logur.LoggerFacade, db *database.DB) dbImpl {
	return dbImpl{
		user:    user.New(logger, db),
		product: product.New(logger, db),
		order:   order.New(logger, db),
	}
}

type dbImpl struct {
	user    user.Repo
	product product.Repo
	order   order.Repo
}

func (i dbImpl) User() user.Repo {
	return i.user
}

func (i dbImpl) Product() product.Repo {
	return i.product
}

func (i dbImpl) Order() order.Repo {
	return i.order
}
