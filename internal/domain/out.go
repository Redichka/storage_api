package domain

import (
	"time"

	"github.com/google/uuid"
)

type OutFilter struct {
	Storage_ref *uuid.UUID
	Price       *float64
	PriceMin    *float64
	PriceMax    *float64
	Date        *time.Time
	DateFrom    *time.Time
	DateTo      *time.Time
}
