package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorageRepo interface {
	Add(ctx context.Context, storage *models.Storage) error
	GetByRef(ctx context.Context, ref uuid.UUID) (*models.Storage, error)
	Update(ctx context.Context, ref uuid.UUID, update *domain.StorageFilter) error
	Delete(ctx context.Context, ref uuid.UUID) error
	ListAll(ctx context.Context) ([]models.Storage, error)
	WithTx(tx *gorm.DB) StorageRepo
	Find(ctx context.Context, filter *domain.StorageFilter) (*[]models.Storage, error)
}

type storageRepo struct {
	db *gorm.DB
}

func (r *storageRepo) Add(ctx context.Context, storage *models.Storage) error {
	return r.db.WithContext(ctx).Create(&storage).Error
}

func (r *storageRepo) Delete(ctx context.Context, ref uuid.UUID) error {
	return r.db.WithContext(ctx).Where("storage_ref = ?", ref).Delete(&models.Storage{}).Error
}

func (r *storageRepo) GetByRef(ctx context.Context, ref uuid.UUID) (*models.Storage, error) {
	var s models.Storage
	err := r.db.WithContext(ctx).First(&s, "storage_ref = ?", ref).Error
	return &s, err

}

func (r *storageRepo) ListAll(ctx context.Context) ([]models.Storage, error) {
	var storages []models.Storage
	err := r.db.WithContext(ctx).Find(&storages).Error
	return storages, err

}

func (r *storageRepo) Update(ctx context.Context, ref uuid.UUID, update *domain.StorageFilter) error {
	updateFields := map[string]interface{}{}
	if update.Name != nil {
		updateFields["name"] = *update.Name
	}
	if update.Address != nil {
		updateFields["address"] = *update.Address
	}
	if len(updateFields) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Model(&models.Storage{}).Where("storage_ref = ?", ref).Updates(updateFields).Error
}

func (r *storageRepo) Find(ctx context.Context, filter *domain.StorageFilter) (*[]models.Storage, error) {
	var storages []models.Storage
	query := r.db.WithContext(ctx)
	if filter.Name != nil {
		query = query.Where("name = ?", *filter.Name)
	}
	if filter.Address != nil {
		query = query.Where("address = ?", *filter.Address)
	}
	err := query.Find(&storages).Error
	return &storages, err
}

func NewStorageRepo(db *gorm.DB) StorageRepo {
	return &storageRepo{db: db}
}

func (r *storageRepo) WithTx(tx *gorm.DB) StorageRepo {
	return &storageRepo{db: tx}
}
