package services

import (
	"errors"
	"time"

	"meddit/models"
	"meddit/repository"
)

type CounselingMenuService interface {
	CreateCounselingMenu(menu *models.CounselingMenu) error
	GetCounselingMenuByID(id uint) (*models.CounselingMenu, error)
	GetCounselingMenuByCounselorID(counselorID uint) ([]models.CounselingMenu, error)
	UpdateCounselingMenu(menu *models.CounselingMenu) error
	DeleteCounselingMenu(menu *models.CounselingMenu) error
}

type counselingMenuService struct {
	counselingMenuRepo repository.CounselingMenuRepository
}

func NewCounselingMenuService(counselingMenuRepo repository.CounselingMenuRepository) CounselingMenuService {
	return &counselingMenuService{
		counselingMenuRepo: counselingMenuRepo,
	}
}

func (s *counselingMenuService) CreateCounselingMenu(menu *models.CounselingMenu) error {
	return s.counselingMenuRepo.Create(menu)
}

func (s *counselingMenuService) GetCounselingMenuByID(id uint) (*models.CounselingMenu, error) {
	return s.counselingMenuRepo.GetByID(id)
}

func (s *counselingMenuService) GetCounselingMenuByCounselorID(counselorID uint) ([]models.CounselingMenu, error) {
	return s.counselingMenuRepo.GetByCounselorID(counselorID)
}

func (s *counselingMenuService) UpdateCounselingMenu(menu *models.CounselingMenu) error {
	if menu.ID == 0 {
		return errors.New("menu id is required")
	}
	menu.UpdatedAt = time.Now()
	return s.counselingMenuRepo.Update(menu)
}

func (s *counselingMenuService) DeleteCounselingMenu(menu *models.CounselingMenu) error {
	if menu.ID == 0 {
		return errors.New("menu id is required")
	}
	return s.counselingMenuRepo.Delete(menu)
}
