package services

import (
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"gorm.io/gorm"
)

type Bills interface {
	AddBill(bill *models.Bill) error
	GetBill(id uuid.UUID) (models.Bill, error)
	UpdateBill(bill *models.Bill) (models.Bill, error)
	CreateBill(bill *models.Bill) (models.Bill, error)
}

type BillService struct {
	*gorm.DB
}

func NewBillService(database *gorm.DB) *BillService {
	return &BillService{DB: database}
}

func (bs *BillService) AddBill(bill *models.Bill) error {
	result := bs.DB.Create(bill)

	return result.Error
}

func (bs *BillService) GetBill(id uuid.UUID) (models.Bill, error) {
	var bill models.Bill
	result := bs.DB.Where(&models.Bill{ID: id}).Find(&bill)

	if result.Error != nil {
		return bill, result.Error
	}

	return bill, nil
}

func (bs *BillService) UpdateBill(bill *models.Bill) (models.Bill, error) {
	result := bs.DB.Save(bill)

	if result.Error != nil {
		return *bill, result.Error
	}

	return *bill, nil
}

func (bs *BillService) CreateBill(bill *models.Bill) (models.Bill, error) {
	bill.ID, _ = uuid.NewV4()

	result := bs.DB.Create(bill)

	if result.Error != nil {
		return *bill, result.Error
	}

	return *bill, nil
}
