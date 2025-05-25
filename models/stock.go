package models

import (
	"time"

	"github.com/google/uuid"
)

type Stock struct {
	Ref         uuid.UUID `gorm:"type:uuid;column:ref;primaryKey"`
	Storage_ref uuid.UUID `gorm:"type:uuid;column:storage_ref;primaryKey"`
	Product_ref uuid.UUID `gorm:"type:uuid;column:product_ref;primaryKey"`
	Ty          string    `gorm:"type:char(1);column:type"`
	Count       float64   `gorm:"columt:count"`
	Date        time.Time `gorm:"type:timestamp;column:date"`
}

func (Stock) TableName() string {
	return "stock"
}
