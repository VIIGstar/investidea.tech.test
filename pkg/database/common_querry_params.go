package database

const (
	DefaultLimit = 50
	MaxLimit     = 500
)

type CommonQueryParams struct {
	ID int64 `json:"id"`

	// by time
	CreatedAtFrom int64 `json:"created_at_from"`
	CreatedAtTo   int64 `json:"created_at_to"`

	// pagination
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	SinceId int64 `json:"since_id"`
}

func (params CommonQueryParams) CorrectRequests() CommonQueryParams {
	if params.Limit <= 0 || params.Limit > MaxLimit {
		params.Limit = DefaultLimit
	}

	return params
}
