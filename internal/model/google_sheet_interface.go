package model

import "context"

type ISpreadsheetRepository interface {
	AppendTroubleshoot(ctx context.Context, log *TroubleshootLog) error
}
