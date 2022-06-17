package user

import (
	"errors"
	"gorm.io/gorm"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/auth"
)

func (i impl) Create(u *entities.User) error {
	return i.db.GormDB().Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(entities.User{}).Create(u).Error
		if txErr != nil {
			return txErr
		}

		switch auth.RoleType(u.Role) {
		case auth.BuyerRole:
			buyer := u.ToBuyerModel()
			buyer.UserID = u.ID
			txErr = tx.Model(entities.Buyer{}).Create(&buyer).Error
		case auth.SellerRole:
			seller := u.ToSellerModel()
			seller.UserID = u.ID
			txErr = tx.Model(entities.Seller{}).Create(&seller).Error
		default:
			return errors.New(dtos.InvalidRoleErrReason)
		}

		return txErr
	})
}
