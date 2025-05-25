package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Product_ref uuid.UUID `gorm:"type:uuid;column:product_ref;primaryKey"`
	Name        string    `gorm:"column:name"`
	Price       float64   `gorm:"column:price"`
}

func (Product) TableName() string {
	return "product"
}
