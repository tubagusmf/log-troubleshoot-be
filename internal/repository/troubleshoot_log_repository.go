package repository

import (
	"context"
	"time"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"

	"gorm.io/gorm"
)

type troubleshootLogRepository struct {
	db *gorm.DB
}

func NewTroubleshootLogRepo(db *gorm.DB) model.ITroubleshootLogRepository {
	return &troubleshootLogRepository{db: db}
}

func (r *troubleshootLogRepository) FindAll(ctx context.Context, log model.TroubleshootLog) ([]*model.TroubleshootLog, error) {
	var logs []*model.TroubleshootLog

	query := r.db.WithContext(ctx).
		Model(&model.TroubleshootLog{}).
		Where("deleted_at IS NULL")

	if log.TicketNumber != "" {
		query = query.Where("ticket_number ILIKE ?", "%"+log.TicketNumber+"%")
	}

	if log.Status != "" {
		query = query.Where("status = ?", log.Status)
	}

	if log.UserID != nil {
		query = query.Where("user_id = ?", *log.UserID)
	}

	if log.ProjectID != nil {
		query = query.Where("project_id = ?", *log.ProjectID)
	}

	if err := query.
		Order("created_at DESC").
		Find(&logs).Error; err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *troubleshootLogRepository) FindByID(ctx context.Context, id int64) (*model.TroubleshootLog, error) {
	var log model.TroubleshootLog

	if err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&log).Error; err != nil {
		return nil, err
	}

	return &log, nil
}

func (r *troubleshootLogRepository) Create(ctx context.Context, log model.TroubleshootLog) (*model.TroubleshootLog, error) {
	if err := r.db.WithContext(ctx).Create(&log).Error; err != nil {
		return nil, err
	}

	return &log, nil
}

func (r *troubleshootLogRepository) Update(ctx context.Context, log model.TroubleshootLog) error {
	return r.db.WithContext(ctx).
		Model(&model.TroubleshootLog{}).
		Where("id = ? AND deleted_at IS NULL", log.ID).
		Updates(log).Error
}

func (r *troubleshootLogRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).
		Model(&model.TroubleshootLog{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
