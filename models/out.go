package models

import (
	"time"

	"github.com/google/uuid"
)

type Out struct {
	Out_ref     uuid.UUID `gorm:"type:uuid;column:out_ref;primaryKey"`
	Storage_ref uuid.UUID `gorm:"type:uuid;column:storage_ref"`
	Price       float64   `gorm:"type:numeric;column:price"`
	Date        time.Time `gorm:"type:timestamp;column:date"`
}

func (Out) TableName() string {
	return "out"
}

type Out_goods struct {
	Out_ref     uuid.UUID `gorm:"type:uuid;column:out_ref;primaryKey"`
	Product_ref uuid.UUID `gorm:"type:uuid;column:product_ref;primaryKey"`
	Quantity    int       `gorm:"column:quantity"`
	Price       float64   `gorm:"type:numeric;column:price"`
}

func (Out_goods) TableName() string {
	return "out_goods"
}
