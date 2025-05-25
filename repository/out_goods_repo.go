package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"gorm.io/gorm"
)

type OutGoodsRepo interface {
	Add(ctx context.Context, out *models.Out_goods) error
	GetByRef(ctx context.Context, key domain.OutGoodsKey) (*models.Out_goods, error)
	Update(ctx context.Context, key domain.OutGoodsKey, update *domain.OutGoodsFilter) error
	Delete(ctx context.Context, key domain.OutGoodsKey) error
	ListAll(ctx context.Context) ([]models.Out_goods, error)
	WithTx(tx *gorm.DB) OutGoodsRepo
	Find(ctx context.Context, filter *domain.OutGoodsFilter) (*[]models.Out_goods, error)
}

type outGoodsRepo struct {
	db *gorm.DB
}

func (r *outGoodsRepo) Add(ctx context.Context, out_goods *models.Out_goods) error {
	return r.db.WithContext(ctx).Create(&out_goods).Error
}

func (r *outGoodsRepo) Delete(ctx context.Context, key domain.OutGoodsKey) error {
	return r.db.WithContext(ctx).Where("out_ref = ? AND product_ref = ?", key.Out_ref, key.Product_ref).Delete(&models.Out_goods{}).Error
}

func (r *outGoodsRepo) GetByRef(ctx context.Context, key domain.OutGoodsKey) (*models.Out_goods, error) {
	var og models.Out_goods
	err := r.db.WithContext(ctx).First(&og, "out_ref = ? AND product_ref = ?", key.Out_ref, key.Product_ref).Error
	return &og, err

}

func (r *outGoodsRepo) ListAll(ctx context.Context) ([]models.Out_goods, error) {
	var ogs []models.Out_goods
	err := r.db.WithContext(ctx).Find(&ogs).Error
	return ogs, err

}

func (r *outGoodsRepo) Update(ctx context.Context, key domain.OutGoodsKey, update *domain.OutGoodsFilter) error {
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
	return r.db.WithContext(ctx).Model(&models.Inc{}).Where("out_ref = ? AND product_ref = ?", key.Out_ref, key.Product_ref).Updates(updateFields).Error
}

func (r *outGoodsRepo) Find(ctx context.Context, filter *domain.OutGoodsFilter) (*[]models.Out_goods, error) {
	var out_goods []models.Out_goods
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
	err := query.Find(&out_goods).Error
	return &out_goods, err
}

func NewOutGoodsRepo(db *gorm.DB) OutGoodsRepo {
	return &outGoodsRepo{db: db}
}

func (r *outGoodsRepo) WithTx(tx *gorm.DB) OutGoodsRepo {
	return &outGoodsRepo{db: tx}
}
