package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type DeviceUsecase struct {
	deviceRepo model.IDeviceRepository
}

func NewDeviceUsecase(deviceRepo model.IDeviceRepository) model.IDeviceUsecase {
	return &DeviceUsecase{
		deviceRepo: deviceRepo,
	}
}

func (d *DeviceUsecase) FindAll(ctx context.Context, device model.Device) ([]*model.Device, error) {
	log := logrus.WithFields(logrus.Fields{
		"filter": device,
	})

	devices, err := d.deviceRepo.FindAll(ctx, device)
	if err != nil {
		log.Error("Failed to fetch devices: ", err)
		return nil, err
	}

	return devices, nil
}

func (d *DeviceUsecase) FindByID(ctx context.Context, id int64) (*model.Device, error) {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	device, err := d.deviceRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch device by ID: ", err)
		return nil, err
	}

	if device == nil {
		return nil, errors.New("device not found")
	}

	return device, nil
}

func (d *DeviceUsecase) Create(ctx context.Context, in model.CreateDeviceInput) (*model.Device, error) {
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

	device := model.Device{
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdDevice, err := d.deviceRepo.Create(ctx, device)
	if err != nil {
		log.Error("Failed to create device: ", err)
		return nil, err
	}

	if createdDevice == nil {
		log.Error("Device not created")
		return nil, errors.New("device not created")
	}

	return createdDevice, nil
}

func (d *DeviceUsecase) Update(ctx context.Context, id int64, in model.UpdateDeviceInput) error {
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

	existingDevice, err := d.deviceRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch device: ", err)
		return err
	}

	if existingDevice == nil || (existingDevice.DeletedAt != nil && !existingDevice.DeletedAt.IsZero()) {
		log.Error("Device is deleted or does not exist")
		return errors.New("device is deleted or does not exist")
	}

	device := model.Device{
		Id:        id,
		Name:      in.Name,
		UpdatedAt: time.Now(),
	}

	err = d.deviceRepo.Update(ctx, device)
	if err != nil {
		log.Error("Failed to update device: ", err)
		return err
	}

	return nil
}

func (d *DeviceUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	if !isAdmin(ctx) {
		return errors.New("forbidden: admin only")
	}

	device, err := d.deviceRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch device for deletion: ", err)
		return err
	}

	if device == nil {
		log.Error("Device not found")
		return errors.New("device not found")
	}

	err = d.deviceRepo.Delete(ctx, id)
	if err != nil {
		log.Error("Failed to delete device: ", err)
		return err
	}

	log.Info("Successfully deleted device with ID: ", id)
	return nil
}
