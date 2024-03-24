package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestPaydayService_Get(t *testing.T) {
	t.Run("Gets paydays associated to user", func(t *testing.T) {
		db, mock, _ := setupDbMock()

		userId := "testUser"
		paydaysWant := []models.Payday{
			{UserID: userId, PayDateOfMonth: 1},
			{UserID: userId, PayDateOfMonth: 15},
		}

		rows := sqlmock.NewRows([]string{"UserID", "PayDateOfMonth"})
		for _, payday := range paydaysWant {
			rows.AddRow(payday.UserID, payday.PayDateOfMonth)
		}
		mock.ExpectQuery(`WHERE "paydays"."user_id" = ?`).WithArgs(userId).WillReturnRows(rows)

		payDayService := NewPayDayService(db)
		payDayService.GetPaydays(userId)
	})
}

func setupDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDbConnection, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDbConnection,
		DriverName: "postgres",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})

	return db, mock, err
}
