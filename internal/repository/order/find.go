package order

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"investidea.tech.test/internal/entities"
	query_params "investidea.tech.test/internal/query-params"
	"investidea.tech.test/pkg/database"
)

func (i impl) Find(ctx context.Context, req query_params.OrderParams, lock bool) ([]entities.Order, error) {
	req.CommonQueryParams = req.CommonQueryParams.CorrectRequests()
	qb := i.getCommonFilter(ctx, req, lock).
		Select(database.SelectColumns(entities.Product{}, database.WithIgnoreFields("seller_id", "buyer_id"))).
		Limit(req.Limit).
		Offset(req.Offset)
	result := make([]entities.Order, 0)
	err := qb.Find(&result).Error
	return result, err
}

func (i impl) getCommonFilter(
	ctx context.Context,
	req query_params.OrderParams,
	lock bool) *gorm.DB {
	qb := i.db.GormDB().WithContext(ctx).Model(entities.Product{})

	if lock {
		qb = qb.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}

	return qb
}
