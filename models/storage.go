package models

import "github.com/google/uuid"

type Storage struct {
	Storage_ref uuid.UUID `gorm:"type:uuid;column:storage_ref;primaryKey"`
	Name        string    `gorm:"column:name"`
	Address     string    `gorm:"column:address"`
}

func (Storage) TableName() string {
	return "storage"
}
