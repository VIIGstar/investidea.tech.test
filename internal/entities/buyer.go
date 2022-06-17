package entities

import base_entity "investidea.tech.test/pkg/base-entity"

type Buyer struct {
	base_entity.Base
	Name             string `json:"name" gorm:"type:varchar(256)"`
	Password         string `json:"password" gorm:"type:varchar(64)"`
	AlamatPengiriman string `json:"alamat_pengiriman" gorm:"type:varchar(1024)"`
	UserID           uint   `json:"user_id"`
}

func (Buyer) TableName() string {
	return "buyers"
}
