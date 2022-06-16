package product

import "investidea.tech.test/internal/entities"

func (i impl) Create(p *entities.Product) error {
	return i.db.GormDB().Model(p).Create(p).Error
}
