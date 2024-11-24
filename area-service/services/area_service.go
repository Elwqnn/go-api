package services

import (
	"area-service/models"
	"area-service/repositories"
	"errors"
)

type AreaService interface {
	CreateArea(area *models.Area) error
	GetAreaByID(id uint) (*models.Area, error)
	GetAllAreas() ([]models.Area, error)
	DeleteArea(id uint) error
}

type areaService struct {
	repo repositories.AreaRepository
}

func NewAreaService(repo repositories.AreaRepository) AreaService {
	return &areaService{repo: repo}
}

func (s *areaService) CreateArea(area *models.Area) error {
	if len(area.Actions) == 0 || len(area.Reactions) == 0 {
		return errors.New("an Area must have at least one Action and one Reaction")
	}
	return s.repo.CreateArea(area)
}

func (s *areaService) GetAreaByID(id uint) (*models.Area, error) {
	return s.repo.GetAreaByID(id)
}

func (s *areaService) GetAllAreas() ([]models.Area, error) {
	return s.repo.GetAllAreas()
}

func (s *areaService) DeleteArea(id uint) error {
	return s.repo.DeleteArea(id)
}
