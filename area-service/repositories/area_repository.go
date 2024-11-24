package repositories

import (
	"area-service/models"

	"gorm.io/gorm"
)

type AreaRepository interface {
	CreateArea(area *models.Area) error
	GetAreaByID(id uint) (*models.Area, error)
	GetAllAreas() ([]models.Area, error)
	DeleteArea(id uint) error
}

type areaRepository struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) AreaRepository {
	return &areaRepository{db: db}
}

func (r *areaRepository) CreateArea(area *models.Area) error {
	return r.db.Create(area).Error
}

func (r *areaRepository) GetAreaByID(id uint) (*models.Area, error) {
	var area models.Area
	err := r.db.Preload("Actions").Preload("Reactions").First(&area, id).Error
	return &area, err
}

func (r *areaRepository) GetAllAreas() ([]models.Area, error) {
	var areas []models.Area
	err := r.db.Preload("Actions").Preload("Reactions").Find(&areas).Error
	return areas, err
}

func (r *areaRepository) DeleteArea(id uint) error {
	return r.db.Delete(&models.Area{}, id).Error
}
