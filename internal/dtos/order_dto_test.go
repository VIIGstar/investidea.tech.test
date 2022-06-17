package dtos

import (
	"github.com/stretchr/testify/assert"
	"investidea.tech.test/internal/entities"
	"testing"
)

func TestOrderDTO_Validate_EmptySellerID(t *testing.T) {
	dto := OrderDTO{
		Order:    entities.Order{},
		Items:    []int{1, 2},
		Quantity: []int{1, 2},
	}
	assert.Equal(t, dto.Validate()[0], UnidentifiedStoreErrText)
	assert.Equal(t, len(dto.Validate()), 1)
}

func TestOrderDTO_Validate_ZeroQuantity(t *testing.T) {
	dto := OrderDTO{
		Order: entities.Order{
			SellerID: 1,
		},
		Items:    []int{1, 2},
		Quantity: []int{0, 0},
	}
	assert.Equal(t, dto.Validate()[0], QuantityItemZeroErrText)
	assert.Equal(t, len(dto.Validate()), 1)
}

func TestOrderDTO_Validate_EmptyItem(t *testing.T) {
	dto := OrderDTO{
		Order: entities.Order{
			SellerID: 1,
		},
		Items:    []int{},
		Quantity: []int{10, 0},
	}
	assert.Equal(t, dto.Validate()[0], EmptyCartErrText)
	assert.Equal(t, len(dto.Validate()), 1)
}

func TestOrderDTO_ToEntity(t *testing.T) {
	dto := OrderDTO{
		Order: entities.Order{
			SellerID: 1,
		},
		Items:    []int{1, 2, 0, 3},
		Quantity: []int{1, 2, 0, 3, 4, 5},
	}
	assert.Nil(t, dto.Validate())
	entity, _ := dto.ToEntity()
	order, ok := entity.(entities.Order)
	assert.True(t, ok)
	assert.Greater(t, len(order.Items), 0)
	assert.Greater(t, len(order.Quantity), 0)
}

func TestOrderDTO_ToEntity2(t *testing.T) {
	dto := OrderDTO{
		Order: entities.Order{
			SellerID: 1,
		},
		Items:    []int{1, 2, 3, 0},
		Quantity: []int{1, 2, 3, 0, 4, 5},
	}
	assert.Nil(t, dto.Validate())
	entity, _ := dto.ToEntity()
	order, ok := entity.(entities.Order)
	assert.True(t, ok)
	assert.Greater(t, len(order.Items), 0)
	assert.Greater(t, len(order.Quantity), 0)
}

func TestOrderDTO_ToEntity3(t *testing.T) {
	dto := OrderDTO{
		Order: entities.Order{
			SellerID: 1,
		},
		Items:    []int{0, 1, 2, 3, 0},
		Quantity: []int{0, 1, 2, 3, 0, 4, 5},
	}
	assert.Nil(t, dto.Validate())
	entity, _ := dto.ToEntity()
	order, ok := entity.(entities.Order)
	assert.True(t, ok)
	assert.Greater(t, len(order.Items), 0)
	assert.Greater(t, len(order.Quantity), 0)
}
