package seeds

import (
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	//payday1, _ := uuid.FromString("bce9b411-319a-447c-bbca-66e16ff54bb5")
	return []Seed{
		//{
		//	Name: "Create Karl",
		//	Run: func(db *gorm.DB) error {
		//		return CreateUser(db, "Karl Merecido", "google-oauth2|110280176458668985185")
		//	},
		//},
		//{
		//	Name: "Create Payday",
		//	Run: func(db *gorm.DB) error {
		//		return db.Create(&models.Payday{ID: payday1, PayDateOfMonth: 1, UserID: "google-oauth2|110280176458668985185",
		//			Bills: []models.Bill{
		//				{
		//					ID:                     uuid.Must(uuid.NewV4()),
		//					DueDateOfMonth:         28,
		//					Name:                   "Car",
		//					AutoPayment:            true,
		//					AutoPaymentDateOfMonth: 15,
		//					Amount:                 decimal.NewFromFloat(500.70),
		//				},
		//			}}).Error
		//	},
		//},
	}
}

func CreateUser(db *gorm.DB, name string, id string) error {
	return db.Create(&models.User{Name: name, ID: id}).Error
}
