package models

import "github.com/gofrs/uuid"

type Payday struct {
	ID             uuid.UUID `json:"id"`
	PayDateOfMonth int       `json:"pay_date_of_month"`
	UserID         string    `json:"user_id"`
	Bills          []Bill    `json:"bills"`
}
