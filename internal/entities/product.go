package entities

type Product struct {
	Base
	ProductName string `json:"product_name" gorm:"type:varchar(256);unique" bind:"required"`
	Description string `json:"description" gorm:"type:text"`
	Price       int64  `json:"price"`
	SellerID    int64  `json:"seller_id;omitempty"`
}

func (Product) TableName() string {
	return "products"
}
