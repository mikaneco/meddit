package repository

import (
	"errors"

	"gorm.io/gorm"

	"meddit/models"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	GetByID(userID uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]models.User, error)
	Delete(user *models.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepo) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepo) GetByID(userID uint) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) Delete(user *models.User) error {
	return repo.db.Delete(user).Error
}
