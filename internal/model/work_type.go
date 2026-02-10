package model

import (
	"context"
	"time"
)

type WorkType struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type CreateWorkTypeInput struct {
	Name string `json:"name" validate:"required,max=100"`
}

type UpdateWorkTypeInput struct {
	Name string `json:"name" validate:"required,max=100"`
}

type IWorkTypeRepository interface {
	FindAll(ctx context.Context, workType WorkType) ([]*WorkType, error)
	FindByID(ctx context.Context, id int64) (*WorkType, error)
	Create(ctx context.Context, workType WorkType) (*WorkType, error)
	Update(ctx context.Context, workType WorkType) error
	Delete(ctx context.Context, id int64) error
}

type IWorkTypeUsecase interface {
	FindAll(ctx context.Context, workType WorkType) ([]*WorkType, error)
	FindByID(ctx context.Context, id int64) (*WorkType, error)
	Create(ctx context.Context, in CreateWorkTypeInput) (*WorkType, error)
	Update(ctx context.Context, id int64, in UpdateWorkTypeInput) error
	Delete(ctx context.Context, id int64) error
}
