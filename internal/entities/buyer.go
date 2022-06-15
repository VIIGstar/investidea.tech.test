package entities

import base_entity "investidea.tech.test/pkg/base-entity"

type Buyer struct {
	base_entity.Base
	Name             string `json:"name" gorm:"type:varchar"`
	Password         string `json:"password" gorm:"type:varchar;size:32"`
	AlamatPengiriman string `json:"alamat_pengiriman" gorm:"type:varchar"`
}
