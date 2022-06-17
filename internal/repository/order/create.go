package order

import (
	"encoding/json"
	"errors"
	"investidea.tech.test/internal/entities"
	info_log "investidea.tech.test/pkg/info-log"
)

func (i impl) Create(order *entities.Order) error {
	var product []entities.Product
	var productIds []int64
	err := json.Unmarshal([]byte(order.Items), &productIds)
	if err != nil {
		return errors.New("invalid request")
	}
	err = i.db.GormDB().
		Model(entities.Product{}).
		Select("id").
		Where("id in ?", productIds).
		Where("seller_id = ?", order.SellerID).
		Find(&product).
		Error
	if err != nil {
		i.logger.Error("query product failed", info_log.ErrorToLogFields("details", err))
		return errors.New("products not found")
	}
	if len(productIds) != len(product) {
		i.logger.Error("product missing")
		return errors.New("some products missing")
	}

	return i.db.GormDB().Model(order).Create(order).Error
}
