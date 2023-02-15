package repository

import (
	"meddit/models"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	CreateAppointment(appointment *models.Appointment) error
	GetAppointmentByID(id uint) (*models.Appointment, error)
	GetAppointmentByUserID(userID uint) ([]models.Appointment, error)
	GetAppointmentByCounselorID(counselorID uint) ([]models.Appointment, error)
	UpdateAppointment(appointment *models.Appointment) error
	CancelAppointment(appointment *models.Appointment) error
	GetOverlappingAppointments(appointment *models.Appointment) ([]models.Appointment, error)
}

type AppointmentRepo struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &AppointmentRepo{
		db: db,
	}
}

func (repo *AppointmentRepo) CreateAppointment(appointment *models.Appointment) error {
	result := repo.db.Create(appointment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *AppointmentRepo) GetAppointmentByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	result := repo.db.Preload("User").Preload("Counselor").Preload("CounselingMenu").Where("id = ?", id).First(&appointment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &appointment, nil
}

func (repo *AppointmentRepo) GetAppointmentByUserID(userID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	result := repo.db.Preload("User").Preload("Counselor").Preload("CounselingMenu").Where("user_id = ?", userID).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}

func (repo *AppointmentRepo) GetAppointmentByCounselorID(counselorID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	result := repo.db.Preload("User").Preload("Counselor").Preload("CounselingMenu").Where("counselor_id = ?", counselorID).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}

func (repo *AppointmentRepo) UpdateAppointment(appointment *models.Appointment) error {
	result := repo.db.Save(appointment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *AppointmentRepo) CancelAppointment(appointment *models.Appointment) error {
	appointment.PaymentState = models.PaymentStateCanceled
	result := repo.db.Save(appointment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *AppointmentRepo) GetOverlappingAppointments(appointment *models.Appointment) ([]models.Appointment, error) {
	var appointments []models.Appointment
	result := repo.db.Where("counselor_id = ?", appointment.CounselorID).Where("start_time <= ?", appointment.EndAt).Where("end_time >= ?", appointment.StartAt).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}
