package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"gorm.io/gorm"
)

type StockRepo interface {
	Add(ctx context.Context, stock *models.Stock) error
	GetByRef(ctx context.Context, key *domain.StockKey) (*models.Stock, error)
	Update(ctx context.Context, key *domain.StockKey, update *domain.StockFilter) error
	Delete(ctx context.Context, key *domain.StockKey) error
	ListAll(ctx context.Context) ([]models.Stock, error)
	WithTx(tx *gorm.DB) StockRepo
	Find(ctx context.Context, filter *domain.StockFilter) (*[]models.Stock, error)
}

type stockRepo struct {
	db *gorm.DB
}

func (r *stockRepo) Add(ctx context.Context, stock *models.Stock) error {
	return r.db.WithContext(ctx).Create(&stock).Error
}

func (r *stockRepo) Delete(ctx context.Context, key *domain.StockKey) error {
	return r.db.WithContext(ctx).Where("ref = ? AND storage_ref = ? AND product_ref = ?", *key.Stock_ref, *key.Storage_ref, *key.Product_ref).Delete(&models.Product{}).Error
}

func (r *stockRepo) GetByRef(ctx context.Context, key *domain.StockKey) (*models.Stock, error) {
	var s models.Stock
	err := r.db.WithContext(ctx).First(&s, "ref = ? AND storage_ref = ? AND product_ref = ?", *key.Stock_ref, *key.Stock_ref, *key.Product_ref).Error
	return &s, err

}

func (r *stockRepo) ListAll(ctx context.Context) ([]models.Stock, error) {
	var stocks []models.Stock
	err := r.db.WithContext(ctx).Find(&stocks).Error
	return stocks, err

}

func (r *stockRepo) Update(ctx context.Context, key *domain.StockKey, update *domain.StockFilter) error {
	updateFields := map[string]interface{}{}
	if update.Ty != nil {
		updateFields["ty"] = *update.Ty
	}
	if update.Count != nil {
		updateFields["count"] = *update.Count
	}
	if update.Count != nil {
		updateFields["date"] = *update.Date
	}
	if len(updateFields) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&models.Stock{}).Where("ref = ? AND storage_ref = ? AND product_ref = ?", *key.Stock_ref, *key.Storage_ref, *key.Product_ref).Updates(updateFields).Error
}

func (r *stockRepo) Find(ctx context.Context, filter *domain.StockFilter) (*[]models.Stock, error) {
	var stocks []models.Stock
	query := r.db.WithContext(ctx)
	if filter.Ty != nil {
		query = query.Where("ty = ?", *filter.Ty)
	}
	if filter.CountFrom != nil {
		query = query.Where("count >= ?", *filter.CountFrom)
	}
	if filter.CountTo != nil {
		query = query.Where("count <= ?", *filter.CountTo)
	}
	if filter.DateFrom != nil {
		query = query.Where("date >= ?", *filter.DateFrom)
	}
	if filter.DateTo != nil {
		query = query.Where("date <= ?", *filter.DateTo)
	}
	err := query.Find(&stocks).Error
	return &stocks, err
}

func NewStockRepo(db *gorm.DB) StockRepo {
	return &stockRepo{db: db}
}

func (r *stockRepo) WithTx(tx *gorm.DB) StockRepo {
	return &stockRepo{db: tx}
}
