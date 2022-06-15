package query_params

import "investidea.tech.test/pkg/database"

type GetUserParams struct {
	database.CommonQueryParams
	Address string `json:"address"`
}
