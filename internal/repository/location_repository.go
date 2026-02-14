package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
	"gorm.io/gorm"
)

type LocationRepo struct {
	db *gorm.DB
}

func NewLocationRepo(db *gorm.DB) model.ILocationRepository {
	return &LocationRepo{
		db: db,
	}
}

func (l *LocationRepo) FindAll(ctx context.Context, location model.Location) ([]*model.Location, error) {
	var locations []*model.Location

	query := l.db.WithContext(ctx).
		Model(&model.Location{}).
		Where("deleted_at IS NULL")

	if location.Name != "" {
		query = query.Where("name ILIKE ?", "%"+location.Name+"%")
	}

	if location.CodeName != "" {
		query = query.Where("code_name ILIKE ?", "%"+location.CodeName+"%")
	}

	err := query.Find(&locations).Error
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (l *LocationRepo) FindByID(ctx context.Context, id int64) (*model.Location, error) {
	var location model.Location

	err := l.db.WithContext(ctx).First(&location, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("location not found")
	}
	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (l *LocationRepo) Create(ctx context.Context, location model.Location) (*model.Location, error) {
	location.CreatedAt = time.Now()
	location.UpdatedAt = time.Now()

	err := l.db.WithContext(ctx).Create(&location).Error
	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (l *LocationRepo) Update(ctx context.Context, location model.Location) error {
	location.UpdatedAt = time.Now()

	err := l.db.WithContext(ctx).
		Model(&model.Location{}).
		Where("id = ? AND deleted_at IS NULL", location.Id).
		Updates(&location).Error
	if err != nil {
		return err
	}

	return nil
}

func (l *LocationRepo) Delete(ctx context.Context, id int64) error {
	err := l.db.WithContext(ctx).
		Model(&model.Location{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *LocationRepo) FindByName(ctx context.Context, name string) (*model.Location, error) {
	var location model.Location

	err := r.db.WithContext(ctx).
		Where("LOWER(name) = LOWER(?) AND deleted_at IS NULL", name).
		First(&location).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("location not found")
	}

	return &location, err
}
