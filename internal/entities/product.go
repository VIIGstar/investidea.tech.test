package entities

import base_entity "investidea.tech.test/pkg/base-entity"

type Product struct {
	base_entity.Base
	ProductName string `json:"product_name" gorm:"type:varchar(256)" bind:"required"`
	Description string `json:"description" gorm:"type:text"`
	Price       int64  `json:"price"`
	SellerID    int64  `json:"seller_id"`
}
