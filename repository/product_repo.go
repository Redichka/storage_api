package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Add(ctx context.Context, product *models.Product) error
	GetByRef(ctx context.Context, ref uuid.UUID) (*models.Product, error)
	Update(ctx context.Context, ref uuid.UUID, update *domain.ProductFilter) error
	Delete(ctx context.Context, ref uuid.UUID) error
	ListAll(ctx context.Context) ([]models.Product, error)
	WithTx(tx *gorm.DB) ProductRepo
	Find(ctx context.Context, key *domain.ProductFilter) (*[]models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func (r *productRepo) Add(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(&product).Error
}

func (r *productRepo) Delete(ctx context.Context, ref uuid.UUID) error {
	return r.db.WithContext(ctx).Where("product_ref = ?", ref).Delete(&models.Product{}).Error
}

func (r *productRepo) GetByRef(ctx context.Context, ref uuid.UUID) (*models.Product, error) {
	var p models.Product
	err := r.db.WithContext(ctx).First(&p, "product_ref = ?", ref).Error
	return &p, err

}

func (r *productRepo) ListAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err

}

func (r *productRepo) Update(ctx context.Context, ref uuid.UUID, update *domain.ProductFilter) error {
	updateFields := map[string]interface{}{}
	if update.Name != nil {
		updateFields["name"] = *update.Name
	}
	if update.Price != nil {
		updateFields["price"] = *update.Price
	}
	if len(updateFields) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&models.Product{}).Where("product_ref = ?", ref).Updates(updateFields).Error
}

func (r *productRepo) Find(ctx context.Context, key *domain.ProductFilter) (*[]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx)
	if key.Name != nil {
		query = query.Where("name = ?", *key.Name)
	}
	if key.PriceFrom != nil {
		query = query.Where("price >= ?", *key.PriceFrom)
	}
	if key.PriceTo != nil {
		query = query.Where("price <= ?", *key.PriceTo)
	}
	err := query.Find(&products).Error
	return &products, err
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) WithTx(tx *gorm.DB) ProductRepo {
	return &productRepo{db: tx}
}
