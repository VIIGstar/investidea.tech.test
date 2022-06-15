package user

import (
	"errors"
	"gorm.io/gorm"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
)

func (i impl) Create(u *entities.User) error {
	return i.db.GormDB().Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(entities.User{}).Create(u).Error
		if txErr != nil {
			return txErr
		}

		switch entities.RoleType(u.Role) {
		case entities.BuyerRole:
			buyer := u.ToBuyerModel()
			txErr = tx.Model(entities.Buyer{}).Create(&buyer).Error
		case entities.SellerRole:
			seller := u.ToSellerModel()
			txErr = tx.Model(entities.Seller{}).Create(&seller).Error
		default:
			return errors.New(dtos.InvalidRoleErrReason)
		}

		return txErr
	})
}
