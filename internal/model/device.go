package model

import (
	"context"
	"time"
)

type Device struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type CreateDeviceInput struct {
	Name string `json:"name" validate:"required,max=100"`
}

type UpdateDeviceInput struct {
	Name string `json:"name" validate:"required,max=100"`
}

type IDeviceRepository interface {
	FindAll(ctx context.Context, device Device) ([]*Device, error)
	FindByID(ctx context.Context, id int64) (*Device, error)
	Create(ctx context.Context, device Device) (*Device, error)
	Update(ctx context.Context, device Device) error
	Delete(ctx context.Context, id int64) error
}

type IDeviceUsecase interface {
	FindAll(ctx context.Context, device Device) ([]*Device, error)
	FindByID(ctx context.Context, id int64) (*Device, error)
	Create(ctx context.Context, in CreateDeviceInput) (*Device, error)
	Update(ctx context.Context, id int64, in UpdateDeviceInput) error
	Delete(ctx context.Context, id int64) error
}
