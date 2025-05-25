package domain

import "github.com/google/uuid"

type IncGoodsFilter struct {
	Quantity     *int
	QuantityFrom *int
	QuantityTo   *int
	Price        *float64
	PriceMin     *float64
	PriceMax     *float64
}

type IncGoodsKey struct {
	Inc_ref     *uuid.UUID
	Product_ref *uuid.UUID
}
