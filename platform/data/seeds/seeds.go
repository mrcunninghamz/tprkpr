package seeds

import (
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	payday1, _ := uuid.FromString("bce9b411-319a-447c-bbca-66e16ff54bb5")
	bill1, _ := uuid.FromString("b3bb7f58-a38c-4c7f-9149-609c61c7c153")
	worksheet1, _ := uuid.FromString("5e926288-8b8e-44ab-9414-40374b675d8e")
	worsheetItem1, _ := uuid.FromString("f57a794d-3c57-4f5b-a989-48613a0804a0")
	worsheetItem2, _ := uuid.FromString("ca50608d-b57f-4989-b486-a9e97bf695d7")
	return []Seed{
		{
			Name: "Upsert Karl",
			Run: func(db *gorm.DB) error {
				return UpsertUser(db, "Karl Merecido", "google-oauth2|110280176458668985185")
			},
		},
		{
			Name: "Upsert Payday",
			Run: func(db *gorm.DB) error {
				return UpsertPayday(db, payday1, []models.Bill{
					{
						ID:                     bill1,
						DueDateOfMonth:         28,
						Name:                   "Car",
						IsAutoPayment:          true,
						AutoPaymentDateOfMonth: 15,
						Amount:                 decimal.NewFromFloat(500.70),
					},
				}, "google-oauth2|110280176458668985185", 1)
			},
		},
		{
			Name: "Upsert worksheet",
			Run: func(db *gorm.DB) error {
				return UpsertWorksheet(db, worksheet1, []models.WorksheetItem{
					{
						ID:             worsheetItem1,
						DueDateOfMonth: 28,
						Name:           "Car",
						Amount:         decimal.NewFromFloat(500.70),
						Notes:          "AutoPayment set to 28 of month",
					},
					{
						ID:             worsheetItem2,
						DueDateOfMonth: 15,
						Name:           "Car2",
						Amount:         decimal.NewFromFloat(700.75),
						Notes:          "AutoPayment set to 15 of month",
						IsPaid:         true,
					},
				}, "google-oauth2|110280176458668985185", 1, 4, 2024, decimal.NewFromFloat(1500.75))
			},
		},
	}
}

func UpsertUser(db *gorm.DB, name string, id string) error {
	user := models.User{Name: name, ID: id}
	if err := db.Where(models.User{ID: id}).Assign(models.User{Name: name}).FirstOrCreate(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpsertPayday(db *gorm.DB, id uuid.UUID, bills []models.Bill, userID string, payDateOfMonth int) error {
	payday := models.Payday{
		ID:             id,
		PayDateOfMonth: payDateOfMonth,
		UserID:         userID,
		Bills:          bills,
	}
	if err := db.Where(models.Payday{ID: id}).Assign(payday).FirstOrCreate(&payday).Error; err != nil {
		return err
	}
	return nil
}

func UpsertWorksheet(db *gorm.DB, id uuid.UUID, items []models.WorksheetItem, userID string, day int, month int, year int, inbank decimal.Decimal) error {
	worksheet := models.Worksheet{
		ID:     id,
		UserID: userID,
		Day:    day,
		Month:  month,
		Year:   year,
		Items:  items,
		InBank: inbank,
	}
	if err := db.Where(models.Worksheet{ID: id}).Assign(worksheet).FirstOrCreate(&worksheet).Error; err != nil {
		return err
	}
	return nil
}
