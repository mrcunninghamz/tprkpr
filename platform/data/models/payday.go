package models

import "github.com/gofrs/uuid"

type Payday struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID         string    `json:"user_id"`
	PayDateOfMonth int       `json:"pay_date_of_month"`
	Bills          []Bill    `json:"bills" gorm:"constraint:OnDelete:CASCADE;"`
}
