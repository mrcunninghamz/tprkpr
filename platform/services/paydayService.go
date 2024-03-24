package services

import (
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"gorm.io/gorm"
)

type Paydays interface {
	Get(userId string) []models.Payday
}

type PaydayService struct {
	*gorm.DB
}

func NewPayDayService(db *gorm.DB) *PaydayService {
	return &PaydayService{
		db,
	}
}

func (service *PaydayService) Get(userId string) []models.Payday {
	var paydays []models.Payday
	service.DB.Preload("Bills").Where(&models.Payday{UserID: userId}).Find(&paydays)

	return paydays
}
