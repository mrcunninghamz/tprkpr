package services

import (
	"github.com/mrcunninghamz/tprkpr/platform/data"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
)

type PaydayService struct {
	DbContext *data.DataContext
}

type Paydays interface {
	Get(userId string) []models.Payday
}

func NewPayDayService(dbContext *data.DataContext) *PaydayService {
	return &PaydayService{
		dbContext,
	}
}

func (service *PaydayService) Get(userId string) []models.Payday {
	var paydays []models.Payday
	service.DbContext.DB.Preload("Bills").Where(&models.Payday{UserID: userId}).Find(&paydays)

	return paydays
}
