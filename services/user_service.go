package services

import (
	"errors"
	"meddit/models"
	"meddit/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	RegisterUser(user *models.User) error
	UpdateUser(user *models.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) RegisterUser(user *models.User) error {
	// 既に同じEmailが登録されていないかチェックする
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if existingUser != nil {
		return errors.New("email is already registered")
	}

	// パスワードをハッシュ化する
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// ユーザーを作成する
	return s.userRepo.Create(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	// パスワードをハッシュ化する
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Update(user)
}
