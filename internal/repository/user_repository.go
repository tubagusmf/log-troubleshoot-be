package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) model.IUserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(ctx context.Context, user model.User) (newUser *model.User, err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindAll(ctx context.Context, user model.User) ([]*model.User, error) {
	var users []*model.User
	query := u.db.WithContext(ctx).Model(&model.User{}).Where("deleted_at IS NULL")

	if user.Username != "" {
		query = query.Where("username LIKE ?", "%"+user.Username+"%")
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) Update(ctx context.Context, user model.User) error {
	user.UpdatedAt = time.Now()

	err := u.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ? AND deleted_at IS NULL", user.Id).
		Updates(user).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) Delete(ctx context.Context, id int64) error {
	err := u.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) FindByCodeName(ctx context.Context, codeName string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Where("code_name = ? AND deleted_at IS NULL", codeName).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
