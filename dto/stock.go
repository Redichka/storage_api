package dto

import (
	"time"

	"github.com/google/uuid"
)

type StockDTO struct {
	Ty        *string    `json:"Ty"`
	Count     *float64   `json:"Count"`
	CountFrom *float64   `json:"CountFrom"`
	CountTo   *float64   `json:"CountTo"`
	Date      *time.Time `json:"Date"`
	DateFrom  *time.Time `json:"DateFrom"`
	DateTo    *time.Time `json:"DateTo"`
}

type StockKeyDTO struct {
	Stock_ref   *uuid.UUID `json:"Stock_ref"`
	Storage_ref *uuid.UUID `json:"Storage_ref"`
	Product_ref *uuid.UUID `json:"Product_ref"`
}
