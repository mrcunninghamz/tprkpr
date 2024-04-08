package models

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type WorksheetItem struct {
	ID             uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name           string          `json:"name"`
	Amount         decimal.Decimal `json:"amount" gorm:"type:numeric(6,2);"`
	IsPaid         bool            `json:"is_paid"`
	DueDateOfMonth int             `json:"due_date_of_month"`
	Notes          string          `json:"notes"`
	WorksheetId    uuid.UUID       `json:"worksheet_id"`
}
