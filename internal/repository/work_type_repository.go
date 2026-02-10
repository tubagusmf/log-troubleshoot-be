package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
	"gorm.io/gorm"
)

type WorkTypeRepo struct {
	db *gorm.DB
}

func NewWorkTypeRepo(db *gorm.DB) model.IWorkTypeRepository {
	return &WorkTypeRepo{
		db: db,
	}
}

func (w *WorkTypeRepo) FindAll(ctx context.Context, workType model.WorkType) ([]*model.WorkType, error) {
	var workTypes []*model.WorkType

	query := w.db.WithContext(ctx).
		Model(&model.WorkType{}).
		Where("deleted_at IS NULL")

	if workType.Name != "" {
		query = query.Where("name ILIKE ?", "%"+workType.Name+"%")
	}

	if err := query.Find(&workTypes).Error; err != nil {
		return nil, err
	}

	return workTypes, nil
}

func (w *WorkTypeRepo) FindByID(ctx context.Context, id int64) (*model.WorkType, error) {
	var workType model.WorkType

	err := w.db.WithContext(ctx).First(&workType, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("work type not found")
	}
	if err != nil {
		return nil, err
	}

	return &workType, nil
}

func (w *WorkTypeRepo) Create(ctx context.Context, workType model.WorkType) (*model.WorkType, error) {
	workType.CreatedAt = time.Now()
	workType.UpdatedAt = time.Now()

	if err := w.db.WithContext(ctx).Create(&workType).Error; err != nil {
		return nil, err
	}

	return &workType, nil
}

func (w *WorkTypeRepo) Update(ctx context.Context, workType model.WorkType) error {
	workType.UpdatedAt = time.Now()

	err := w.db.WithContext(ctx).Model(&model.WorkType{}).Where("id = ? AND deleted_at IS NULL", workType.Id).Updates(&workType).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkTypeRepo) Delete(ctx context.Context, id int64) error {
	err := w.db.WithContext(ctx).Model(&model.WorkType{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
