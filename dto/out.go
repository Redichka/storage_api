package dto

import (
	"time"

	"github.com/google/uuid"
)

type OutFilter struct {
	Storage_ref *uuid.UUID `json:"Storage_ref"`
	PriceMin    *float64   `json:"PriceMin"`
	PriceMax    *float64   `json:"PriceMax"`
	DateFrom    *time.Time `json:"DateFrom"`
	DateTo      *time.Time `json:"DateTo"`
}
