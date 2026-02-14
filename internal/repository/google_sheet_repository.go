package repository

import (
	"context"
	"fmt"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetRepository struct {
	service       *sheets.Service
	spreadsheetID string
	sheetName     string
}

func NewGoogleSheetRepository(
	credentialFile string,
	spreadsheetID string,
	sheetName string,
) (*GoogleSheetRepository, error) {

	ctx := context.Background()

	srv, err := sheets.NewService(ctx,
		option.WithCredentialsFile(credentialFile),
		option.WithScopes(sheets.SpreadsheetsScope),
	)
	if err != nil {
		return nil, err
	}

	return &GoogleSheetRepository{
		service:       srv,
		spreadsheetID: spreadsheetID,
		sheetName:     sheetName,
	}, nil
}

func (r *GoogleSheetRepository) AppendTroubleshoot(
	ctx context.Context,
	log *model.TroubleshootLog,
) error {

	values := [][]interface{}{
		{
			log.TroubleDate.Format("2006-01-02"),
			log.TroubleTime.Format("15:04:05"),
			log.Part,
			log.DeviceNumber,
			log.Issue,
			log.Status,
			log.WhatsappSender,
		},
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := r.service.Spreadsheets.Values.Append(
		r.spreadsheetID,
		fmt.Sprintf("%s!A:G", r.sheetName),
		valueRange,
	).ValueInputOption("RAW").Do()

	return err
}
