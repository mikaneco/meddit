package repository

import (
	"errors"

	"gorm.io/gorm"

	"meddit/models"
)

type CounselorRepository interface {
	Create(counselor *models.Counselor) error
	Update(counselor *models.Counselor) error
	GetByID(counselorID uint) (*models.Counselor, error)
	GetByEmail(email string) (*models.Counselor, error)
	GetAll() ([]models.Counselor, error)
	Delete(counselor *models.Counselor) error
}

type CounselorRepo struct {
	db *gorm.DB
}

func NewCounselorRepository(db *gorm.DB) CounselorRepository {
	return &CounselorRepo{db: db}
}

func (repo *CounselorRepo) Create(counselor *models.Counselor) error {
	return repo.db.Create(counselor).Error
}

func (repo *CounselorRepo) Update(counselor *models.Counselor) error {
	return repo.db.Save(counselor).Error
}

func (repo *CounselorRepo) GetByID(counselorID uint) (*models.Counselor, error) {
	var counselor models.Counselor
	if err := repo.db.First(&counselor, counselorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &counselor, nil
}

func (repo *CounselorRepo) GetByEmail(email string) (*models.Counselor, error) {
	var counselor models.Counselor
	if err := repo.db.Where("email = ?", email).First(&counselor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &counselor, nil
}

func (repo *CounselorRepo) GetAll() ([]models.Counselor, error) {
	var counselors []models.Counselor
	if err := repo.db.Find(&counselors).Error; err != nil {
		return nil, err
	}
	return counselors, nil
}

func (repo *CounselorRepo) Delete(counselor *models.Counselor) error {
	return repo.db.Delete(counselor).Error
}
