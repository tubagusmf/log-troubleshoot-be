package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tubagusmf/log-troubleshoot-be/internal/helper"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type ParsedWhatsAppMessage struct {
	Project  string
	Location string
	Part     string
	DeviceID string
	Issue    string
	CodeName string
}

type WhatsAppConsumerUsecase struct {
	userRepo         model.IUserRepository
	projectRepo      model.IProjectRepository
	locationRepo     model.ILocationRepository
	troubleshootRepo model.ITroubleshootLogRepository
	sheetRepo        model.ISpreadsheetRepository
}

func NewWhatsAppConsumerUsecase(
	userRepo model.IUserRepository,
	troubleshootRepo model.ITroubleshootLogRepository,
	sheetRepo model.ISpreadsheetRepository,
) *WhatsAppConsumerUsecase {
	return &WhatsAppConsumerUsecase{
		userRepo:         userRepo,
		troubleshootRepo: troubleshootRepo,
		sheetRepo:        sheetRepo,
	}
}

func parseMessage(message string) ParsedWhatsAppMessage {
	lines := strings.Split(message, "\n")

	result := ParsedWhatsAppMessage{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Project :") {
			result.Project = strings.TrimSpace(strings.TrimPrefix(line, "Project :"))
		}
		if strings.HasPrefix(line, "Stasiun :") {
			result.Location = strings.TrimSpace(strings.TrimPrefix(line, "Stasiun :"))
		}
		if strings.HasPrefix(line, "Part :") {
			result.Part = strings.TrimSpace(strings.TrimPrefix(line, "Part :"))
		}
		if strings.HasPrefix(line, "ID :") {
			result.DeviceID = strings.TrimSpace(strings.TrimPrefix(line, "ID :"))
		}
		if strings.HasPrefix(line, "Permasalahan :") {
			result.Issue = strings.TrimSpace(strings.TrimPrefix(line, "Permasalahan :"))
		}
		if strings.HasPrefix(line, "#") {
			result.CodeName = strings.TrimPrefix(line, "#")
		}
	}

	return result
}

func (u *WhatsAppConsumerUsecase) Consume(
	ctx context.Context,
	payload model.WhatsAppWebhookRequest,
) error {

	parsed := helper.ParseWhatsAppReport(payload.Message)

	if parsed.Project == "" || parsed.Issue == "" {
		return errors.New("invalid report format")
	}

	user, err := u.userRepo.FindByCodeName(ctx, parsed.CodeName)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}
	if user == nil {
		return errors.New("user is nil")
	}

	project, err := u.projectRepo.FindByName(ctx, parsed.Project)
	if err != nil {
		return fmt.Errorf("project not found: %w", err)
	}
	if project == nil {
		return errors.New("project is nil")
	}

	location, err := u.locationRepo.FindByName(ctx, parsed.Station)
	if err != nil {
		return fmt.Errorf("location not found: %w", err)
	}
	if location == nil {
		return errors.New("location is nil")
	}

	log := model.TroubleshootLog{
		UserID:          &user.Id,
		ProjectID:       &project.Id,
		LocationID:      &location.Id,
		DeviceNumber:    parsed.DeviceID,
		Part:            parsed.Part,
		Issue:           parsed.Issue,
		Status:          "OPEN",
		WhatsappSender:  payload.Sender,
		WhatsappMessage: payload.Message,
	}

	fmt.Println("User:", user)
	fmt.Println("Project:", project)
	fmt.Println("Location:", location)

	_, err = u.troubleshootRepo.Create(ctx, log)
	return err
}
