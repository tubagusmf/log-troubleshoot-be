package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type ProjectUsecase struct {
	projectRepo model.IProjectRepository
}

func NewProjectUsecase(projectRepo model.IProjectRepository) model.IProjectUsecase {
	return &ProjectUsecase{
		projectRepo: projectRepo,
	}
}

func (p *ProjectUsecase) FindAll(ctx context.Context, project model.Project) ([]*model.Project, error) {
	log := logrus.WithFields(logrus.Fields{
		"filter": project,
	})

	projects, err := p.projectRepo.FindAll(ctx, project)
	if err != nil {
		log.Error("Failed to fetch projects: ", err)
		return nil, err
	}

	return projects, nil
}

func (p *ProjectUsecase) FindByID(ctx context.Context, id int64) (*model.Project, error) {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	project, err := p.projectRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch project by ID: ", err)
		return nil, err
	}

	if project == nil {
		log.Error("Project not found")
		return nil, errors.New("project not found")
	}

	return project, nil
}

func (p *ProjectUsecase) Create(ctx context.Context, in model.CreateProjectInput) (*model.Project, error) {
	log := logrus.WithFields(logrus.Fields{
		"in": in,
	})

	if !isAdmin(ctx) {
		return nil, errors.New("forbidden: admin only")
	}

	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error("Validation error:", err)
		return nil, err
	}

	project := model.Project{
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdProject, err := p.projectRepo.Create(ctx, project)
	if err != nil {
		log.Error("Failed to create project: ", err)
		return nil, err
	}

	if createdProject == nil {
		log.Error("Project not created")
		return nil, errors.New("project not created")
	}

	return createdProject, nil
}

func (p *ProjectUsecase) Update(ctx context.Context, id int64, in model.UpdateProjectInput) error {
	log := logrus.WithFields(logrus.Fields{
		"in": in,
	})

	if !isAdmin(ctx) {
		return errors.New("forbidden: admin only")
	}

	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error("Validation error:", err)
		return err
	}

	existingProject, err := p.projectRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch project: ", err)
		return err
	}

	if existingProject == nil || (existingProject.DeletedAt != nil && !existingProject.DeletedAt.IsZero()) {
		log.Error("Project is deleted or does not exist")
		return errors.New("project is deleted or does not exist")
	}

	project := model.Project{
		Id:        id,
		Name:      in.Name,
		UpdatedAt: time.Now(),
	}

	err = p.projectRepo.Update(ctx, project)
	if err != nil {
		log.Error("Failed to update project: ", err)
		return err
	}

	return nil
}

func (p *ProjectUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	if !isAdmin(ctx) {
		return errors.New("forbidden: admin only")
	}

	project, err := p.projectRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to fetch project for deletion: ", err)
		return err
	}

	if project == nil {
		log.Error("Project not found")
		return errors.New("project not found")
	}

	err = p.projectRepo.Delete(ctx, id)
	if err != nil {
		log.Error("Failed to delete project: ", err)
		return err
	}

	log.Info("Successfully deleted project with ID: ", id)

	return nil
}
