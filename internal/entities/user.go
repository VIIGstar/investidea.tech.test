package entities

import base_entity "investidea.tech.test/pkg/base-entity"

type User struct {
	base_entity.Base
	Role     string `json:"role" gorm:"type:varchar(16)"`
	Email    string `json:"email" gorm:"type:varchar(256);unique" bind:"required"`
	Name     string `json:"name" gorm:"-"`
	Password string `json:"password" gorm:"-" bind:"required"`
	Address  string `json:"address" gorm:"-"`
}

func (u User) ToBuyerModel() Buyer {
	return Buyer{
		Name:             u.Name,
		Password:         u.Password,
		AlamatPengiriman: u.Address,
	}
}

func (u User) ToSellerModel() Seller {
	return Seller{
		Name:         u.Name,
		Password:     u.Password,
		AlamatPickup: u.Address,
	}
}

func (User) TableName() string {
	return "users"
}
