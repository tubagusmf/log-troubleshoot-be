package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
	"gorm.io/gorm"
)

type DeviceRepo struct {
	db *gorm.DB
}

func NewDeviceRepo(db *gorm.DB) model.IDeviceRepository {
	return &DeviceRepo{
		db: db,
	}
}

func (d *DeviceRepo) FindAll(ctx context.Context, device model.Device) ([]*model.Device, error) {
	var devices []*model.Device

	query := d.db.WithContext(ctx).
		Model(&model.Device{}).
		Where("deleted_at IS NULL")

	if device.Name != "" {
		query = query.Where("name LIKE ?", "%"+device.Name+"%")
	}

	err := query.Find(&devices).Error
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (d *DeviceRepo) FindByID(ctx context.Context, id int64) (*model.Device, error) {
	var device model.Device

	err := d.db.WithContext(ctx).First(&device, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("device not found")
	}
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (d *DeviceRepo) Create(ctx context.Context, device model.Device) (*model.Device, error) {
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()

	err := d.db.WithContext(ctx).Create(&device).Error
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (d *DeviceRepo) Update(ctx context.Context, device model.Device) error {
	device.UpdatedAt = time.Now()

	err := d.db.WithContext(ctx).
		Model(&model.Device{}).
		Where("id = ? AND deleted_at IS NULL", device.Id).
		Updates(&device).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *DeviceRepo) Delete(ctx context.Context, id int64) error {
	err := d.db.WithContext(ctx).Model(&model.Device{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}
