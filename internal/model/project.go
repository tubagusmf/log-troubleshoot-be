package model

import (
	"context"
	"time"
)

type Project struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type CreateProjectInput struct {
	Name string `json:"name" validate:"required,max=100"`
}

type UpdateProjectInput struct {
	Name string `json:"name" validate:"required,max=100"`
}

type IProjectRepository interface {
	Create(ctx context.Context, project Project) (*Project, error)
	FindAll(ctx context.Context, project Project) ([]*Project, error)
	FindByID(ctx context.Context, id int64) (*Project, error)
	Update(ctx context.Context, project Project) error
	Delete(ctx context.Context, id int64) error
	FindByName(ctx context.Context, name string) (*Project, error)
}

type IProjectUsecase interface {
	FindAll(ctx context.Context, project Project) ([]*Project, error)
	FindByID(ctx context.Context, id int64) (*Project, error)
	Create(ctx context.Context, in CreateProjectInput) (*Project, error)
	Update(ctx context.Context, id int64, in UpdateProjectInput) error
	Delete(ctx context.Context, id int64) error
}
