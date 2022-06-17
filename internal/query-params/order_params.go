package query_params

import "investidea.tech.test/pkg/database"

type OrderParams struct {
	database.CommonQueryParams
	Status   string `json:"status"`
	SellerID int64  `json:"seller_id"`
}
