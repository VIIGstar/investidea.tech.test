package product

import (
	"context"
	"investidea.tech.test/internal/entities"
	query_params "investidea.tech.test/internal/query-params"
)

func (i impl) Get(ctx context.Context, req query_params.ProductParams, lock bool) (*entities.Product, error) {
	req.CommonQueryParams = req.CommonQueryParams.CorrectRequests()
	qb := i.getCommonFilter(ctx, req, lock)

	result := &entities.Product{}
	err := qb.First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
