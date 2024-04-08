package services

import (
	"github.com/gofrs/uuid"
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"gorm.io/gorm"
)

type Worksheets interface {
	CreateWorksheet(worksheet *models.Worksheet) (*models.Worksheet, error)
	GetWorksheets(userId string) ([]models.Worksheet, error)
	GetWorksheet(id uuid.UUID) (models.Worksheet, error)
	UpdateWorksheet(worksheet *models.Worksheet) error
	DeleteWorksheet(worksheetId uuid.UUID) error
	GetWorksheetItem(id uuid.UUID) (models.WorksheetItem, error)
	UpdateWorksheetItem(worksheetItem *models.WorksheetItem) error
}

type WorksheetService struct {
	*gorm.DB
}

func NewWorksheetService(db *gorm.DB) *WorksheetService {
	return &WorksheetService{
		db,
	}
}

func (service *WorksheetService) GetWorksheets(userId string) ([]models.Worksheet, error) {
	var worksheets []models.Worksheet

	result := service.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		db = db.Order("due_date_of_month asc")
		return db
	}).Where(&models.Worksheet{UserID: userId}).Find(&worksheets)
	if result.Error != nil {
		return nil, result.Error
	}

	return worksheets, nil
}

func (service *WorksheetService) CreateWorksheet(worksheet *models.Worksheet) (*models.Worksheet, error) {
	result := service.DB.Create(worksheet)
	if result.Error != nil {
		return nil, result.Error
	}

	return worksheet, nil
}

func (service *WorksheetService) GetWorksheet(id uuid.UUID) (models.Worksheet, error) {
	var worksheet models.Worksheet
	result := service.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		db = db.Order("due_date_of_month asc")
		return db
	}).First(&worksheet, id)
	return worksheet, result.Error
}

func (service *WorksheetService) UpdateWorksheet(worksheet *models.Worksheet) error {
	result := service.DB.Save(*worksheet)
	return result.Error
}

func (service *WorksheetService) DeleteWorksheet(worksheetId uuid.UUID) error {
	var worksheet models.Worksheet

	result := service.DB.Where(&models.Worksheet{ID: worksheetId}).Delete(&worksheet)
	return result.Error
}

func (service *WorksheetService) GetWorksheetItem(id uuid.UUID) (models.WorksheetItem, error) {
	var worksheetItem models.WorksheetItem
	result := service.DB.First(&worksheetItem, id)
	return worksheetItem, result.Error
}

func (service *WorksheetService) UpdateWorksheetItem(worksheetItem *models.WorksheetItem) error {
	result := service.DB.Save(*worksheetItem)
	return result.Error
}
