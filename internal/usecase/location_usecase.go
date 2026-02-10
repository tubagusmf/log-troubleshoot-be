package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type LocationUsecase struct {
	locationRepo model.ILocationRepository
}

func NewLocationUsecase(locationRepo model.ILocationRepository) model.ILocationUsecase {
	return &LocationUsecase{
		locationRepo: locationRepo,
	}
}

func (l *LocationUsecase) FindAll(ctx context.Context, location model.Location) ([]*model.Location, error) {
	log := logrus.WithFields(logrus.Fields{
		"filter": location,
	})

	locations, err := l.locationRepo.FindAll(ctx, location)
	if err != nil {
		log.Error("Failed to fetch locations: ", err)
		return nil, err
	}

	return locations, nil
}

func (l *LocationUsecase) FindByID(ctx context.Context, id int64) (*model.Location, error) {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	location, err := l.locationRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch location by ID: ", err)
		return nil, err
	}

	if location == nil {
		log.Error("Location not found")
		return nil, errors.New("location not found")
	}

	return location, nil
}

func (l *LocationUsecase) Create(ctx context.Context, in model.CreateLocationInput) (*model.Location, error) {
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

	location := model.Location{
		Name:      in.Name,
		CodeName:  in.CodeName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdLocation, err := l.locationRepo.Create(ctx, location)
	if err != nil {
		log.Error("Failed to create location: ", err)
		return nil, err
	}

	if createdLocation == nil {
		log.Error("Location not created")
		return nil, errors.New("location not created")
	}

	return createdLocation, nil
}

func (l *LocationUsecase) Update(ctx context.Context, id int64, in model.UpdateLocationInput) error {
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

	existingLocation, err := l.locationRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if existingLocation == nil || (existingLocation.DeletedAt != nil && !existingLocation.DeletedAt.IsZero()) {
		log.Error("Location is deleted or does not exist")
		return errors.New("location is deleted or does not exist")
	}

	location := model.Location{
		Id:        id,
		Name:      in.Name,
		CodeName:  in.CodeName,
		UpdatedAt: time.Now(),
	}

	err = l.locationRepo.Update(ctx, location)
	if err != nil {
		return err
	}

	return nil
}

func (l *LocationUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	if !isAdmin(ctx) {
		return errors.New("forbidden: admin only")
	}

	location, err := l.locationRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch location for deletion: ", err)
		return err
	}

	if location == nil {
		log.Error("Location not found")
		return errors.New("location not found")
	}

	err = l.locationRepo.Delete(ctx, id)
	if err != nil {
		log.Error("Failed to delete location: ", err)
		return err
	}

	log.Info("Successfully deleted location with ID: ", id)
	return nil
}
