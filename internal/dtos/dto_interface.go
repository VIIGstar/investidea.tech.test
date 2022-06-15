package dtos

type IEntityTransformer interface {
	ToEntity() (interface{}, error)
	Validate() []string
}
