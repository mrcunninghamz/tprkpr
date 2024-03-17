package data

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

var TprKprContext DataContext

type DataContext struct {
	Paydays []Payday
	Bills   []Bill
}

type Payday struct {
	Id             uuid.UUID
	PayDateOfMonth int
}

type Bill struct {
	Id                     uuid.UUID
	DueDateOfMonth         int
	Name                   string
	Notes                  string
	AutoPayment            bool
	AutoPaymentDateOfMonth int
	Amount                 decimal.Decimal
	PaydayId               uuid.UUID
}

func init() {
	payday1 := uuid.Must(uuid.NewV4())
	payday15 := uuid.Must(uuid.NewV4())

	TprKprContext.Paydays = []Payday{
		{
			Id:             payday1,
			PayDateOfMonth: 1,
		},
		{
			Id:             payday15,
			PayDateOfMonth: 15,
		},
	}
	TprKprContext.Bills = []Bill{
		{
			Id:             uuid.Must(uuid.NewV4()),
			Name:           "Car",
			DueDateOfMonth: 8,
			Amount:         decimal.NewFromFloat(500.70),
			PaydayId:       payday1,
		},
		{
			Id:             uuid.Must(uuid.NewV4()),
			Name:           "Electricity",
			DueDateOfMonth: 17,
			Amount:         decimal.NewFromFloat(100.75),
			PaydayId:       payday15,
		},
	}
}
