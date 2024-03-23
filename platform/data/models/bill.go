package models

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Bill struct {
	ID                     uuid.UUID
	DueDateOfMonth         int
	Name                   string
	Notes                  string
	AutoPayment            bool
	AutoPaymentDateOfMonth int
	Amount                 decimal.Decimal
	PaydayId               uuid.UUID
}
