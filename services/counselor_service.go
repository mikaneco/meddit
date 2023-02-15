package services

import (
	"errors"

	"meddit/models"
	"meddit/repository"
	"meddit/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CounselorService interface {
	RegisterCounselor(counselor *models.Counselor) (*models.Counselor, error)
	Login(email, password string) (*models.Counselor, error)
	GetCounselorByID(id uint) (*models.Counselor, error)
	UpdateCounselor(counselor *models.Counselor) error
}

type counselorService struct {
	counselorRepo      repository.CounselorRepository
	counselingMenuRepo repository.CounselingMenuRepository
	availableDayRepo   repository.AvailableDayRepository
	appointmentRepo    repository.AppointmentRepository
}

func NewCounselorService(counselorRepo repository.CounselorRepository, counselingMenuRepo repository.CounselingMenuRepository, availableDayRepo repository.AvailableDayRepository, appointmentRepo repository.AppointmentRepository) CounselorService {
	return &counselorService{
		counselorRepo:      counselorRepo,
		counselingMenuRepo: counselingMenuRepo,
		availableDayRepo:   availableDayRepo,
		appointmentRepo:    appointmentRepo,
	}
}

func (s *counselorService) RegisterCounselor(counselor *models.Counselor) (*models.Counselor, error) {
	// 既に同じEmailが登録されていないかチェックする
	existingCounselor, err := s.counselorRepo.GetByEmail(counselor.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingCounselor != nil {
		return nil, errors.New("email is already registered")
	}

	// パスワードをハッシュ化する
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(counselor.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	counselor.Password = string(hashedPassword)

	// カウンセラーを作成する
	err = s.counselorRepo.Create(counselor)
	if err != nil {
		return nil, err
	}

	// カウンセリングメニューを作成する
	defaultMenu := models.CounselingMenu{
		Name:        "標準メニュー",
		Description: "カウンセリング内容を記載してください",
		CounselorID: counselor.ID,
		Price:       75000,
		Duration:    60,
	}

	if err := s.counselingMenuRepo.Create(&defaultMenu); err != nil {
		return nil, err
	}

	return counselor, nil
}

func (s *counselorService) Login(email, password string) (*models.Counselor, error) {
	counselor, err := s.counselorRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, counselor.Password) {
		return nil, errors.New("invalid email or password")
	}

	return counselor, nil
}

func (s *counselorService) GetCounselorByID(id uint) (*models.Counselor, error) {
	return s.counselorRepo.GetByID(id)
}

func (s *counselorService) UpdateCounselor(counselor *models.Counselor) error {
	return s.counselorRepo.Update(counselor)
}
