package usecase

import (
	"context"
	"errors"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type troubleshootLogUsecase struct {
	repo model.ITroubleshootLogRepository
}

func NewTroubleshootLogUsecase(repo model.ITroubleshootLogRepository) model.ITroubleshootLogUsecase {
	return &troubleshootLogUsecase{repo: repo}
}

func (u *troubleshootLogUsecase) FindAll(ctx context.Context, log model.TroubleshootLog) ([]*model.TroubleshootLog, error) {
	return u.repo.FindAll(ctx, log)
}

func (u *troubleshootLogUsecase) FindByID(ctx context.Context, id int64) (*model.TroubleshootLog, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *troubleshootLogUsecase) Create(ctx context.Context, log model.TroubleshootLog) (*model.TroubleshootLog, error) {
	if log.Issue == "" {
		return nil, errors.New("issue is required")
	}

	if log.Status == "" {
		log.Status = "OPEN"
	}

	return u.repo.Create(ctx, log)
}

func (u *troubleshootLogUsecase) Update(ctx context.Context, id int64, log model.TroubleshootLog) error {
	log.ID = id
	return u.repo.Update(ctx, log)
}

func (u *troubleshootLogUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
