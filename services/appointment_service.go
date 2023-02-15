package services

import (
	"time"

	"meddit/errors"
	"meddit/models"
	"meddit/repository"
)

type AppointmentService interface {
	GetAppointmentByID(id uint) (*models.Appointment, error)
	GetUserAppointments(userID uint) ([]models.Appointment, error)
	GetCounselorAppointments(counselorID uint) ([]models.Appointment, error)
	CreateAppointment(appointment *models.Appointment) error
	UpdateAppointment(appointment *models.Appointment) error
	CancelAppointment(appointment *models.Appointment) error
}

type appointmentService struct {
	appointmentRepo  repository.AppointmentRepository
	availableDayRepo repository.AvailableDayRepository
	counselorRepo    repository.CounselorRepository
}

func NewAppointmentService(appointmentRepo repository.AppointmentRepository, availableDayRepo repository.AvailableDayRepository, counselorRepo repository.CounselorRepository) AppointmentService {
	return &appointmentService{
		appointmentRepo:  appointmentRepo,
		availableDayRepo: availableDayRepo,
		counselorRepo:    counselorRepo,
	}
}

func (s *appointmentService) GetAppointmentByID(id uint) (*models.Appointment, error) {
	return s.appointmentRepo.GetAppointmentByID(id)
}

func (s *appointmentService) GetUserAppointments(userID uint) ([]models.Appointment, error) {
	return s.appointmentRepo.GetAppointmentByUserID(userID)
}

func (s *appointmentService) GetCounselorAppointments(counselorID uint) ([]models.Appointment, error) {
	return s.appointmentRepo.GetAppointmentByCounselorID(counselorID)
}

func (s *appointmentService) CreateAppointment(appointment *models.Appointment) error {
	// バリデーションのチェック
	if appointment.StartAt.Before(time.Now()) {
		return errors.ErrInvalidAppointmentTime
	}
	if appointment.CounselorID == 0 {
		return errors.ErrInvalidCounselor
	}
	if appointment.UserID == 0 {
		return errors.ErrInvalidUser
	}
	if appointment.CounselingMenuID == 0 {
		return errors.ErrInvalidCounselingMenu
	}

	// 重複予約のチェック
	existing, err := s.appointmentRepo.GetOverlappingAppointments(appointment)
	if err != nil {
		return err
	}
	if len(existing) > 0 {
		return errors.ErrAppointmentOverlapping
	}

	// CounselorのAvailableTimeのチェック
	if err := s.checkCounselorAvailability(appointment); err != nil {
		return err
	}

	// アポイントメントの作成
	if err := s.appointmentRepo.CreateAppointment(appointment); err != nil {
		return err
	}

	// 決済処理

	return nil
}

func (s *appointmentService) UpdateAppointment(appointment *models.Appointment) error {
	// バリデーションのチェック

	if err := s.appointmentRepo.UpdateAppointment(appointment); err != nil {
		return err
	}

	// 決済処理

	return nil
}

func (s *appointmentService) CancelAppointment(appointment *models.Appointment) error {
	// バリデーションのチェック

	if err := s.appointmentRepo.CancelAppointment(appointment); err != nil {
		return err
	}

	// 決済処理

	return nil
}

func (s *appointmentService) checkCounselorAvailability(appointment *models.Appointment) error {
	// CounselorのAvailableTimeのチェック
	availableDays, err := s.availableDayRepo.GetAllByCounselorID(appointment.CounselorID)
	if err != nil {
		return err
	}
	for _, availableDay := range availableDays {
		if availableDay.IsAvailableAt(appointment.StartAt) {
			return nil
		}
	}
	return errors.ErrCounselorNotAvailable
}
