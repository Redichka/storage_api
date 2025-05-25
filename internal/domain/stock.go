package domain

import (
	"time"

	"github.com/google/uuid"
)

type StockFilter struct {
	Ty        *string
	Count     *float64
	CountFrom *float64
	CountTo   *float64
	Date      *time.Time
	DateFrom  *time.Time
	DateTo    *time.Time
}

type StockKey struct {
	Stock_ref   *uuid.UUID
	Storage_ref *uuid.UUID
	Product_ref *uuid.UUID
}
