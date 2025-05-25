package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"gorm.io/gorm"
)

type IncGoodsRepo interface {
	Add(ctx context.Context, inc_goods *models.Inc_goods) error
	GetByRef(ctx context.Context, key domain.IncGoodsKey) (*models.Inc_goods, error)
	Update(ctx context.Context, key domain.IncGoodsKey, update *domain.IncGoodsFilter) error
	Delete(ctx context.Context, key domain.IncGoodsKey) error
	ListAll(ctx context.Context) ([]models.Inc_goods, error)
	WithTx(tx *gorm.DB) IncGoodsRepo
	Find(ctx context.Context, filter *domain.IncGoodsFilter) (*[]models.Inc_goods, error)
}

type incGoodsRepo struct {
	db *gorm.DB
}

func (r *incGoodsRepo) Add(ctx context.Context, inc_goods *models.Inc_goods) error {
	return r.db.WithContext(ctx).Create(&inc_goods).Error
}

func (r *incGoodsRepo) Delete(ctx context.Context, key domain.IncGoodsKey) error {
	return r.db.WithContext(ctx).Where("inc_ref = ? AND product_ref = ?", key.Inc_ref, key.Product_ref).Delete(&models.Inc_goods{}).Error
}

func (r *incGoodsRepo) GetByRef(ctx context.Context, key domain.IncGoodsKey) (*models.Inc_goods, error) {
	var ig models.Inc_goods
	err := r.db.WithContext(ctx).First(&ig, "inc_ref = ? AND product_ref = ?", key.Inc_ref, key.Product_ref).Error
	return &ig, err

}

func (r *incGoodsRepo) ListAll(ctx context.Context) ([]models.Inc_goods, error) {
	var igs []models.Inc_goods
	err := r.db.WithContext(ctx).Find(&igs).Error
	return igs, err

}

func (r *incGoodsRepo) Update(ctx context.Context, key domain.IncGoodsKey, update *domain.IncGoodsFilter) error {
	updateFields := map[string]interface{}{}
	if update.Price != nil {
		updateFields["price"] = *update.Price
	}
	if update.Quantity != nil {
		updateFields["quantity"] = *update.Quantity
	}
	if len(updateFields) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&models.Inc{}).Where("inc_ref = ? AND product_ref = ?", key.Inc_ref, key.Product_ref).Updates(updateFields).Error
}

func (r *incGoodsRepo) Find(ctx context.Context, filter *domain.IncGoodsFilter) (*[]models.Inc_goods, error) {
	var inc_goods []models.Inc_goods
	query := r.db.WithContext(ctx)
	if filter.PriceMin != nil {
		query = query.Where("price >= ?", *filter.PriceMin)
	}
	if filter.PriceMax != nil {
		query = query.Where("price <= ?", *filter.PriceMax)
	}
	if filter.QuantityFrom != nil {
		query = query.Where("quantity >= ?", *filter.PriceMin)
	}
	if filter.PriceMin != nil {
		query = query.Where("quantity >= ?", *filter.PriceMin)
	}
	err := query.Find(&inc_goods).Error
	return &inc_goods, err
}

func NewIncGoodsRepo(db *gorm.DB) IncGoodsRepo {
	return &incGoodsRepo{db: db}
}

func (r *incGoodsRepo) WithTx(tx *gorm.DB) IncGoodsRepo {
	return &incGoodsRepo{db: tx}
}
