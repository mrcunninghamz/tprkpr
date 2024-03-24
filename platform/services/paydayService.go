package services

import (
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"gorm.io/gorm"
)

type Paydays interface {
	GetPaydays(userId string) []models.Payday
	GetPayday(paydayId uuid.UUID) (models.Payday, error)
	CreatePayday(payday *models.Payday) (models.Payday, error)
}

type PaydayService struct {
	*gorm.DB
}

func NewPayDayService(db *gorm.DB) *PaydayService {
	return &PaydayService{
		db,
	}
}

func (service *PaydayService) GetPaydays(userId string) []models.Payday {
	var paydays []models.Payday
	service.DB.Preload("Bills").Where(&models.Payday{UserID: userId}).Find(&paydays)

	return paydays
}

func (service *PaydayService) GetPayday(paydayId uuid.UUID) (models.Payday, error) {
	var payday models.Payday

	result := service.DB.Preload("Bills").Where(&models.Payday{ID: paydayId}).Find(&payday)
	if result.Error != nil {
		return models.Payday{}, result.Error
	}
	return payday, nil
}

func (service *PaydayService) CreatePayday(payday *models.Payday) (models.Payday, error) {
	payday.ID, _ = uuid.NewV4()

	result := service.DB.Create(payday)
	if result.Error != nil {
		return *payday, result.Error
	}
	return *payday, nil
}
