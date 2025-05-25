package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OutRepo interface {
	Add(ctx context.Context, out *models.Out) error
	GetByRef(ctx context.Context, ref uuid.UUID) (*models.Out, error)
	Update(ctx context.Context, ref uuid.UUID, update *domain.OutFilter) error
	Delete(ctx context.Context, ref uuid.UUID) error
	ListAll(ctx context.Context) ([]models.Out, error)
	WithTx(tx *gorm.DB) OutRepo
	Find(ctx context.Context, filter *domain.OutFilter) (*[]models.Out, error)
}

type outRepo struct {
	db *gorm.DB
}

func (r *outRepo) Add(ctx context.Context, out *models.Out) error {
	return r.db.WithContext(ctx).Create(&out).Error
}

func (r *outRepo) Delete(ctx context.Context, ref uuid.UUID) error {
	return r.db.WithContext(ctx).Where("out_ref = ?", ref).Delete(&models.Out{}).Error
}

func (r *outRepo) GetByRef(ctx context.Context, ref uuid.UUID) (*models.Out, error) {
	var o models.Out
	err := r.db.WithContext(ctx).First(&o, "out_ref = ?", ref).Error
	return &o, err

}

func (r *outRepo) ListAll(ctx context.Context) ([]models.Out, error) {
	var outs []models.Out
	err := r.db.WithContext(ctx).Find(&outs).Error
	return outs, err

}

func (r *outRepo) Update(ctx context.Context, ref uuid.UUID, update *domain.OutFilter) error {
	updateFields := map[string]interface{}{}
	if update.Date != nil {
		updateFields["date"] = *update.Date
	}
	if update.Price != nil {
		updateFields["price"] = *update.Price
	}
	if len(updateFields) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&models.Out{}).Where("out_ref = ?", ref).Updates(updateFields).Error
}

func (r *outRepo) Find(ctx context.Context, filter *domain.OutFilter) (*[]models.Out, error) {
	var outs []models.Out
	query := r.db.WithContext(ctx)
	if filter.PriceMin != nil {
		query = query.Where("count >= ?", *filter.PriceMin)
	}
	if filter.PriceMax != nil {
		query = query.Where("count <= ?", *filter.PriceMax)
	}
	if filter.DateFrom != nil {
		query = query.Where("date >= ?", *filter.DateFrom)
	}
	if filter.DateTo != nil {
		query = query.Where("date <= ?", *filter.DateTo)
	}
	err := query.Find(&outs).Error
	return &outs, err
}

func NewOutRepo(db *gorm.DB) OutRepo {
	return &outRepo{db: db}
}

func (r *outRepo) WithTx(tx *gorm.DB) OutRepo {
	return &outRepo{db: tx}
}
