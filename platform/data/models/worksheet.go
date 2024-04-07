package models

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Worksheet struct {
	ID           uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID       string          `json:"user_id"`
	Month        int             `json:"month"`
	Day          int             `json:"day"`
	Year         int             `json:"year"`
	InBank       decimal.Decimal `json:"inBank" gorm:"type:numeric(6,2);"`
	Total        decimal.Decimal `json:"total" gorm:"type:numeric(6,2);"`
	Net          decimal.Decimal `json:"net" gorm:"->;type:numeric(6,2) GENERATED ALWAYS AS (in_bank - total) STORED;"`
	TotalPaid    decimal.Decimal `json:"totalPaid" gorm:"type:numeric(6,2);"`
	NetAfterPaid decimal.Decimal `json:"netAfterPaid" gorm:"->;type:numeric(6,2) GENERATED ALWAYS AS (in_bank - total_paid) STORED;"`
	Items        []WorksheetItem `json:"items" gorm:"constraint:OnDelete:CASCADE;"`
}
