package model

import (
	"context"
	"time"
)

type TroubleshootLog struct {
	ID              int64      `gorm:"primaryKey" json:"id"`
	TicketNumber    string     `gorm:"unique" json:"ticket_number"`
	TroubleDate     time.Time  `json:"trouble_date"`
	TroubleTime     time.Time  `json:"trouble_time"`
	DoneDate        *time.Time `json:"done_date"`
	DoneTime        *time.Time `json:"done_time"`
	Duration        *string    `json:"duration"`
	UserID          *int64     `json:"user_id"`
	ProjectID       *int64     `json:"project_id"`
	LocationID      *int64     `json:"location_id"`
	DeviceID        *int64     `json:"device_id"`
	WorkTypeID      *int64     `json:"work_type_id"`
	DeviceNumber    string     `json:"device_number"`
	Part            string     `json:"part"`
	Issue           string     `json:"issue"`
	Solution        string     `json:"solution"`
	Status          string     `json:"status"`
	WhatsappSender  string     `json:"whatsapp_sender"`
	WhatsappMessage string     `json:"whatsapp_message"`
	SheetID         string     `json:"sheet_id"`
	SheetRow        *int       `json:"sheet_row"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"-"`
}

type ITroubleshootLogRepository interface {
	FindAll(ctx context.Context, log TroubleshootLog) ([]*TroubleshootLog, error)
	FindByID(ctx context.Context, id int64) (*TroubleshootLog, error)
	Create(ctx context.Context, log TroubleshootLog) (*TroubleshootLog, error)
	Update(ctx context.Context, log TroubleshootLog) error
	Delete(ctx context.Context, id int64) error
}

type ITroubleshootLogUsecase interface {
	FindAll(ctx context.Context, log TroubleshootLog) ([]*TroubleshootLog, error)
	FindByID(ctx context.Context, id int64) (*TroubleshootLog, error)
	Create(ctx context.Context, log TroubleshootLog) (*TroubleshootLog, error)
	Update(ctx context.Context, id int64, log TroubleshootLog) error
	Delete(ctx context.Context, id int64) error
}
