package entities

import base_entity "investidea.tech.test/pkg/base-entity"

type Seller struct {
	base_entity.Base
	Name         string `json:"name" gorm:"type:varchar(256)"`
	Password     string `json:"password" gorm:"type:varchar(64)"`
	AlamatPickup string `json:"alamat_pickup" gorm:"type:varchar(1024)"`
}
