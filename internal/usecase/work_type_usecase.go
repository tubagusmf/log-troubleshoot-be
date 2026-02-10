package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type WorkTypeUsecase struct {
	workTypeRepo model.IWorkTypeRepository
}

func NewWorkTypeUsecase(repo model.IWorkTypeRepository) model.IWorkTypeUsecase {
	return &WorkTypeUsecase{
		workTypeRepo: repo,
	}
}

func (w *WorkTypeUsecase) FindAll(ctx context.Context, workType model.WorkType) ([]*model.WorkType, error) {
	log := logrus.WithFields(logrus.Fields{
		"filter": workType,
	})

	data, err := w.workTypeRepo.FindAll(ctx, workType)
	if err != nil {
		log.Error("Failed to fetch work types: ", err)
		return nil, err
	}

	return data, nil
}

func (w *WorkTypeUsecase) FindByID(ctx context.Context, id int64) (*model.WorkType, error) {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	workType, err := w.workTypeRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch work type: ", err)
		return nil, err
	}

	if workType == nil {
		log.Error("Work type not found")
		return nil, errors.New("work type not found")
	}

	return workType, nil
}

func (w *WorkTypeUsecase) Create(ctx context.Context, in model.CreateWorkTypeInput) (*model.WorkType, error) {
	log := logrus.WithFields(logrus.Fields{
		"in": in,
	})

	if !isAdmin(ctx) {
		return nil, errors.New("forbidden: admin only")
	}

	if err := v.StructCtx(ctx, in); err != nil {
		log.Error("Validation error: ", err)
		return nil, err
	}

	workType := model.WorkType{
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := w.workTypeRepo.Create(ctx, workType)
	if err != nil {
		log.Error("Failed to create work type: ", err)
		return nil, err
	}

	if created == nil {
		log.Error("Work type not created")
		return nil, errors.New("work type not created")
	}

	return created, nil
}

func (w *WorkTypeUsecase) Update(ctx context.Context, id int64, in model.UpdateWorkTypeInput) error {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
		"in": in,
	})

	if !isAdmin(ctx) {
		return errors.New("forbidden: admin only")
	}

	if err := v.StructCtx(ctx, in); err != nil {
		log.Error("Validation error: ", err)
		return err
	}

	existing, err := w.workTypeRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch work type: ", err)
		return err
	}

	if existing == nil || (existing.DeletedAt != nil && !existing.DeletedAt.IsZero()) {
		log.Error("Work type is deleted or does not exist")
		return errors.New("work type is deleted or does not exist")
	}

	workType := model.WorkType{
		Id:        id,
		Name:      in.Name,
		UpdatedAt: time.Now(),
	}

	err = w.workTypeRepo.Update(ctx, workType)
	if err != nil {
		log.Error("Failed to update work type: ", err)
		return err
	}

	return nil
}

func (w *WorkTypeUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	if !isAdmin(ctx) {
		return errors.New("forbidden: admin only")
	}

	workType, err := w.workTypeRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch work type: ", err)
		return err
	}

	if workType == nil {
		log.Error("Work type not found")
		return errors.New("work type not found")
	}

	err = w.workTypeRepo.Delete(ctx, id)
	if err != nil {
		log.Error("Failed to delete work type: ", err)
		return err
	}

	log.Info("Successfully deleted work type with ID: ", id)
	return nil
}
