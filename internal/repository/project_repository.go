package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) model.IProjectRepository {
	return &ProjectRepo{
		db: db,
	}
}

func (p *ProjectRepo) FindAll(ctx context.Context, project model.Project) ([]*model.Project, error) {
	var projects []*model.Project
	query := p.db.WithContext(ctx).Model(&model.Project{}).Where("deleted_at IS NULL")

	if project.Name != "" {
		query = query.Where("name LIKE ?", "%"+project.Name+"%")
	}

	err := query.Find(&projects).Error
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (p *ProjectRepo) FindByID(ctx context.Context, id int64) (*model.Project, error) {
	var project model.Project
	err := p.db.WithContext(ctx).First(&project, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("project not found")
	}
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (p *ProjectRepo) Create(ctx context.Context, project model.Project) (*model.Project, error) {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	err := p.db.WithContext(ctx).Create(&project).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *ProjectRepo) Update(ctx context.Context, project model.Project) error {
	project.UpdatedAt = time.Now()

	err := p.db.WithContext(ctx).Model(&model.Project{}).Where("id = ? AND deleted_at IS NULL", project.Id).Updates(&project).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectRepo) Delete(ctx context.Context, id int64) error {
	err := p.db.WithContext(ctx).Model(&model.Project{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}
