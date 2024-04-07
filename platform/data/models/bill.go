package models

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Bill struct {
	ID                     uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	DueDateOfMonth         int             `json:"due_date_of_month"`
	Name                   string          `json:"name"`
	Notes                  string          `json:"notes"`
	IsAutoPayment          bool            `json:"auto_payment"`
	AutoPaymentDateOfMonth int             `json:"auto_payment_date_of_month"`
	Amount                 decimal.Decimal `json:"amount" gorm:"type:numeric(6,2);"`
	PaydayId               uuid.UUID       `json:"payday_id"`
}
