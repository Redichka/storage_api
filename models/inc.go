package models

import (
	"time"

	"github.com/google/uuid"
)

type Inc struct {
	Inc_ref     uuid.UUID `gorm:"type:uuid;column:inc_ref;primaryKey"`
	Storage_ref uuid.UUID `gorm:"type:uuid;column:storage_ref"`
	Price       float64   `gorm:"type:numeric;column:price"`
	Date        time.Time `gorm:"type:timestamp;column:date"`
}

func (Inc) TableName() string {
	return "inc"
}

type Inc_goods struct {
	Inc_ref     uuid.UUID `gorm:"type:uuid;column:inc_ref;primaryKey"`
	Product_ref uuid.UUID `gorm:"type:uuid;column:product_ref;primaryKey"`
	Quantity    int       `gorm:"column:quantity"`
	Price       float64   `gorm:"type:numeric;column:price"`
}

func (Inc_goods) TableName() string {
	return "inc_goods"
}
