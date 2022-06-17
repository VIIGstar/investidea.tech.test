package order

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"investidea.tech.test/internal/entities"
	query_params "investidea.tech.test/internal/query-params"
	"investidea.tech.test/pkg/database"
)

func (i impl) Update(ctx context.Context,
	value entities.Order,
	conditions query_params.OrderParams,
	lock bool) (*entities.Order, error) {
	qb := i.commonFilterUpdate(ctx, conditions, lock).
		Updates(&value)

	if qb.Error != nil {
		return nil, qb.Error
	}

	if qb.RowsAffected == 0 {
		return nil, database.NoChanges
	}
	return &value, nil
}

func (i impl) commonFilterUpdate(ctx context.Context, conditions query_params.OrderParams, lock bool) *gorm.DB {
	qb := i.db.GormDB().
		WithContext(ctx).
		Model(entities.Order{}).
		Where("id = ?", conditions.ID)

	if lock {
		qb = qb.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}

	if len(conditions.Status) > 0 {
		qb = qb.Where("status = ?", conditions.Status)
	}

	if conditions.SellerID > 0 {
		qb = qb.Where("seller_id = ? ", conditions.SellerID)
	}
	return qb
}
