package models

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Bill struct {
	ID                     uuid.UUID       `json:"id"`
	DueDateOfMonth         int             `json:"due_date_of_month"`
	Name                   string          `json:"name"`
	Notes                  string          `json:"notes"`
	AutoPayment            bool            `json:"auto_payment"`
	AutoPaymentDateOfMonth int             `json:"auto_payment_date_of_month"`
	Amount                 decimal.Decimal `json:"amount"`
	PaydayId               uuid.UUID       `json:"payday_id"`
}
