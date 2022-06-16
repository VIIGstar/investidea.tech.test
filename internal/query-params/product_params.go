package query_params

import "investidea.tech.test/pkg/database"

type ProductParams struct {
	database.CommonQueryParams
	Keyword string `json:"keyword"`
}
