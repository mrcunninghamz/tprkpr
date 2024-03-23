package data

import (
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/platform/data/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DataContext struct {
	*gorm.DB
}

func Connect() (*DataContext, error) {
	connectionString := os.Getenv("POSTGRESSQL_CONNECTION")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {

		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Payday{}, &models.Bill{})
	if err != nil {

		return nil, err
	}

	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	return &DataContext{
		DB: db,
	}, nil
}

//func init() {
//	payday1 := uuid.Must(uuid.NewV4())
//	payday15 := uuid.Must(uuid.NewV4())
//
//	TprKprContext.Paydays = []Payday{
//		{
//			ID:             payday1,
//			PayDateOfMonth: 1,
//		},
//		{
//			ID:             payday15,
//			PayDateOfMonth: 15,
//		},
//	}
//	TprKprContext.Bills = []Bill{
//		{
//			ID:             uuid.Must(uuid.NewV4()),
//			Name:           "Car",
//			DueDateOfMonth: 8,
//			Amount:         decimal.NewFromFloat(500.70),
//			PaydayId:       payday1,
//		},
//		{
//			ID:             uuid.Must(uuid.NewV4()),
//			Name:           "Electricity",
//			DueDateOfMonth: 17,
//			Amount:         decimal.NewFromFloat(100.75),
//			PaydayId:       payday15,
//		},
//	}
//}
