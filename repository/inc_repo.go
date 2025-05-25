package repo

import (
	"context"
	"storage_api/internal/domain"
	"storage_api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IncRepo interface {
	Add(ctx context.Context, inc *models.Inc) error
	GetByRef(ctx context.Context, ref uuid.UUID) (*models.Inc, error)
	Update(ctx context.Context, ref uuid.UUID, update *domain.IncFilter) error
	Delete(ctx context.Context, ref uuid.UUID) error
	ListAll(ctx context.Context) ([]models.Inc, error)
	WithTx(tx *gorm.DB) IncRepo
	Find(ctx context.Context, filter *domain.IncFilter) (*[]models.Inc, error)
}

type incRepo struct {
	db *gorm.DB
}

func (r *incRepo) Add(ctx context.Context, inc *models.Inc) error {
	return r.db.WithContext(ctx).Create(&inc).Error
}

func (r *incRepo) Delete(ctx context.Context, ref uuid.UUID) error {
	return r.db.WithContext(ctx).Where("inc_ref = ?", ref).Delete(&models.Inc{}).Error
}

func (r *incRepo) GetByRef(ctx context.Context, ref uuid.UUID) (*models.Inc, error) {
	var i models.Inc
	err := r.db.WithContext(ctx).First(&i, "inc_ref = ?", ref).Error
	return &i, err

}

func (r *incRepo) ListAll(ctx context.Context) ([]models.Inc, error) {
	var incs []models.Inc
	err := r.db.WithContext(ctx).Find(&incs).Error
	return incs, err

}

func (r *incRepo) Update(ctx context.Context, ref uuid.UUID, update *domain.IncFilter) error {
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
	return r.db.WithContext(ctx).Model(&models.Inc{}).Where("inc_ref = ?", ref).Updates(updateFields).Error
}

func (r *incRepo) Find(ctx context.Context, filter *domain.IncFilter) (*[]models.Inc, error) {
	var incs []models.Inc
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
	err := query.Find(&incs).Error
	return &incs, err
}

func NewIncRepo(db *gorm.DB) IncRepo {
	return &incRepo{db: db}
}

func (r *incRepo) WithTx(tx *gorm.DB) IncRepo {
	return &incRepo{db: tx}
}
