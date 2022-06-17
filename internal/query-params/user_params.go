package query_params

import "investidea.tech.test/pkg/database"

type GetUserParams struct {
	database.CommonQueryParams
	Email    string `json:"email"`
	Password string `json:"password"`
}
