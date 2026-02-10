package model

import (
	"context"
	"time"
)

type Location struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	CodeName  string     `json:"code_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type CreateLocationInput struct {
	Name     string `json:"name" validate:"required,max=100"`
	CodeName string `json:"code_name" validate:"required,max=10"`
}

type UpdateLocationInput struct {
	Name     string `json:"name" validate:"required,max=100"`
	CodeName string `json:"code_name" validate:"required,max=10"`
}

type ILocationRepository interface {
	FindAll(ctx context.Context, location Location) ([]*Location, error)
	FindByID(ctx context.Context, id int64) (*Location, error)
	Create(ctx context.Context, location Location) (*Location, error)
	Update(ctx context.Context, location Location) error
	Delete(ctx context.Context, id int64) error
}

type ILocationUsecase interface {
	FindAll(ctx context.Context, location Location) ([]*Location, error)
	FindByID(ctx context.Context, id int64) (*Location, error)
	Create(ctx context.Context, in CreateLocationInput) (*Location, error)
	Update(ctx context.Context, id int64, in UpdateLocationInput) error
	Delete(ctx context.Context, id int64) error
}
