package entities

import base_entity "investidea.tech.test/pkg/base-entity"

type Order struct {
	base_entity.Base
	BuyerID                    int64    `json:"buyer_id"`
	SellerID                   int64    `json:"seller_id"`
	DeliverySourceAddress      string   `json:"delivery_source_address" gorm:"type:varchar"`
	DeliveryDestinationAddress string   `json:"delivery_destination_address" gorm:"type:varchar"`
	Items                      []string `json:"items"`
	Quantity                   []int64  `json:"quantity"`
	Price                      int64    `json:"price"`
	TotalPrice                 int64    `json:"total_price"`
}
