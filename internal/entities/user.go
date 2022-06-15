package entities

import base_entity "investidea.tech.test/pkg/base-entity"

const (
	BuyerRole  = RoleType("buyer")
	SellerRole = RoleType("seller")
)

type User struct {
	base_entity.Base
	Role     string `json:"role" gorm:"type:varchar;size:16"`
	Email    string `json:"email" gorm:"type:varchar"`
	Name     string `json:"name" gorm:"-"`
	Password string `json:"password" gorm:"-"`
	Address  string `json:"address" gorm:"-"`
}

type RoleType string

func (t RoleType) String() string {
	return string(t)
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
