package dtos

import (
	"encoding/json"
	"investidea.tech.test/internal/entities"
	"math"
)

const (
	EmptyCartErrText         = "Empty cart"
	QuantityItemZeroErrText  = "Quantity item is zero"
	UnidentifiedStoreErrText = "Unidentified store"
)

type OrderDTO struct {
	entities.Order
	Items    []int `json:"items"`
	Quantity []int `json:"quantity"`
}

func (dto OrderDTO) Validate() []string {
	var errs []string

	if len(dto.Items) == 0 || len(dto.Quantity) == 0 {
		errs = append(errs, "Empty cart")
	}

	totalQuantity := 0
	for _, q := range dto.Quantity {
		totalQuantity += q
	}
	if totalQuantity <= 0 {
		errs = append(errs, "Quantity item is zero")
	}

	if dto.SellerID == 0 {
		errs = append(errs, "Unidentified store")
	}

	return errs
}

func (dto OrderDTO) ToEntity() (interface{}, error) {
	order := dto.Order

	minLength := math.Min(float64(len(dto.Items)), float64(len(dto.Quantity)))
	dto.Items = dto.Items[:int(minLength)]
	dto.Quantity = dto.Quantity[:int(minLength)]
	for i := len(dto.Quantity) - 1; i >= 0; i-- {
		if dto.Quantity[i] == 0 {
			//remove this
			dto.Quantity = append(dto.Quantity[:i], dto.Quantity[i+1:]...)
			dto.Items = append(dto.Items[:i], dto.Items[i+1:]...)
		}
	}

	bItem, _ := json.Marshal(dto.Items)
	bQty, _ := json.Marshal(dto.Quantity)

	order.Items = string(bItem)
	order.Quantity = string(bQty)

	return order, nil
}
