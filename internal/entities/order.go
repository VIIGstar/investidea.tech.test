package entities

const (
	OrderStatusPending  = OrderStatusType("pending")
	OrderStatusAccepted = OrderStatusType("accepted")
)

type Order struct {
	Base
	BuyerID int64 `json:"buyer_id"`
	// Should parse from domain store instead of input body
	SellerID                   int64  `json:"seller_id"`
	DeliverySourceAddress      string `json:"delivery_source_address" gorm:"type:text"`
	DeliveryDestinationAddress string `json:"delivery_destination_address" gorm:"type:text"`
	Items                      string `json:"items" gorm:"type:text"`
	Quantity                   string `json:"quantity" gorm:"type:text"`
	Price                      int64  `json:"price"`
	TotalPrice                 int64  `json:"total_price"`
	// Includes: OrderStatusPending | OrderStatusAccepted
	Status string `json:"status" gorm:"type:varchar(32)"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderStatusType string

func (t OrderStatusType) String() string {
	return string(t)
}
