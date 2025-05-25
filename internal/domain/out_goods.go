package domain

import "github.com/google/uuid"

type OutGoodsFilter struct {
	Quantity     *int
	QuantityFrom *int
	QuantityTo   *int
	Price        *float64
	PriceMin     *float64
	PriceMax     *float64
}

type OutGoodsKey struct {
	Out_ref     *uuid.UUID
	Product_ref *uuid.UUID
}
