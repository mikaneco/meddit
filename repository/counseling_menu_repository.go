package repository

import (
	"gorm.io/gorm"

	"meddit/models"
)

type CounselingMenuRepository interface {
	GetByCounselorID(counselorID uint) ([]models.CounselingMenu, error)
	GetByID(id uint) (*models.CounselingMenu, error)
	Create(counselingMenu *models.CounselingMenu) error
	Update(counselingMenu *models.CounselingMenu) error
	Delete(counselingMenu *models.CounselingMenu) error
}

type CounselingMenuRepo struct {
	DB *gorm.DB
}

func NewCounselingMenuRepository(db *gorm.DB) CounselingMenuRepository {
	return &CounselingMenuRepo{DB: db}
}

func (repo *CounselingMenuRepo) GetByCounselorID(counselorID uint) ([]models.CounselingMenu, error) {
	var counselingMenus []models.CounselingMenu
	if err := repo.DB.Where("counselor_id = ?", counselorID).Find(&counselingMenus).Error; err != nil {
		return nil, err
	}
	return counselingMenus, nil
}

func (repo *CounselingMenuRepo) GetByID(id uint) (*models.CounselingMenu, error) {
	var counselingMenu models.CounselingMenu
	if err := repo.DB.Where("id = ?", id).First(&counselingMenu).Error; err != nil {
		return nil, err
	}
	return &counselingMenu, nil
}

func (repo *CounselingMenuRepo) Create(counselingMenu *models.CounselingMenu) error {
	return repo.DB.Create(counselingMenu).Error
}

func (repo *CounselingMenuRepo) Update(counselingMenu *models.CounselingMenu) error {
	return repo.DB.Save(counselingMenu).Error
}

func (repo *CounselingMenuRepo) Delete(counselingMenu *models.CounselingMenu) error {
	return repo.DB.Delete(counselingMenu).Error
}
