package entities

type Seller struct {
	Base
	Name         string `json:"name" gorm:"type:varchar(256)"`
	Password     string `json:"password" gorm:"type:varchar(64)"`
	AlamatPickup string `json:"alamat_pickup" gorm:"type:varchar(1024)"`
	UserID       uint   `json:"user_id"`
}

func (Seller) TableName() string {
	return "sellers"
}
