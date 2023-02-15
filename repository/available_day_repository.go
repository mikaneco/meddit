package repository

import (
	"errors"
	"meddit/models"

	"gorm.io/gorm"
)

type AvailableDayRepository interface {
	GetByID(id uint) (*models.AvailableDay, error)
	GetAllByCounselorID(counselorID uint) ([]models.AvailableDay, error)
	Create(availableDay *models.AvailableDay) error
	Update(availableDay *models.AvailableDay) error
	Delete(availableDay *models.AvailableDay) error
}

type AvailableDayRepo struct {
	DB *gorm.DB
}

func NewAvailableDayRepository(db *gorm.DB) AvailableDayRepository {
	return &AvailableDayRepo{DB: db}
}

func (repo *AvailableDayRepo) GetByID(id uint) (*models.AvailableDay, error) {
	var availableDay models.AvailableDay
	if err := repo.DB.Preload("Counselor").Preload("AvailableTimes").First(&availableDay, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &availableDay, nil
}

func (repo *AvailableDayRepo) GetCounselorAvailableDayByID(counselorID uint, availableDayID uint) (*models.AvailableDay, error) {
	var availableDay models.AvailableDay
	if err := repo.DB.Preload("Counselor").Preload("AvailableTimes").Where("counselor_id = ?", counselorID).First(&availableDay, availableDayID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &availableDay, nil
}


func (repo *AvailableDayRepo) GetAllByCounselorID(counselorID uint) ([]models.AvailableDay, error) {
	var availableDays []models.AvailableDay
	if err := repo.DB.Preload("Counselor").Preload("AvailableTimes").Where("counselor_id = ?", counselorID).Find(&availableDays).Error; err != nil {
		return nil, err
	}
	return availableDays, nil
}

func (repo *AvailableDayRepo) Create(availableDay *models.AvailableDay) error {
	return repo.DB.Create(availableDay).Error
}

func (repo *AvailableDayRepo) Update(availableDay *models.AvailableDay) error {
	return repo.DB.Save(availableDay).Error
}

func (repo *AvailableDayRepo) Delete(availableDay *models.AvailableDay) error {
	return repo.DB.Delete(availableDay).Error
}
