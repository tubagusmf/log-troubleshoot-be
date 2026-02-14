package model

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ContextAuthKey string

const BearerAuthKey ContextAuthKey = "BearerAuth"

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type User struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	CodeName  string     `json:"code_name"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type IUserRepository interface {
	FindAll(ctx context.Context, user User) ([]*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id int64) error
	FindByCodeName(ctx context.Context, codeName string) (*User, error)
}

type IUserUsecase interface {
	FindAll(ctx context.Context, user User) ([]*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)
	Login(ctx context.Context, in LoginInput) (token string, err error)
	Create(ctx context.Context, in CreateUserInput) (token string, err error)
	Update(ctx context.Context, id int64, in UpdateUserInput) error
	Delete(ctx context.Context, id int64) error
}

type LoginInput struct {
	Id       int64  `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,max=100"`
	CodeName string `json:"code_name" validate:"required,max=10"`
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,min=3,max=50"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name" validate:"required,max=100"`
	CodeName string `json:"code_name" validate:"required,max=10"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=3,max=50"`
	Role     string `json:"role" validate:"required"`
}
