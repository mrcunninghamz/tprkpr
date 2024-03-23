package models

import "github.com/gofrs/uuid"

type Payday struct {
	ID             uuid.UUID
	PayDateOfMonth int
	UserID         string
	Bills          []Bill
}
