package database

type ColumnParseOption interface {
	IsOption() bool
}

type IgnoreColumnsOption struct {
	Fields []string
}

func (opt IgnoreColumnsOption) IsOption() bool {
	return true
}

func WithIgnoreFields(fields ...string) ColumnParseOption {
	return IgnoreColumnsOption{
		Fields: fields,
	}
}
