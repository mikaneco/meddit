package services

import (
	"errors"
	"time"

	"meddit/models"
	"meddit/repository"
)

type AvailableDayService interface {
	CreateAvailableDay(day *models.AvailableDay) error
	CreateTmpAvailableDay(day *models.AvailableDay) error
	CreateHoliday(day *models.AvailableDay) error
	UpdateAvailableDay(day *models.AvailableDay) error
	UpdateTmpAvailableDay(day *models.AvailableDay) error
	UpdateHoliday(day *models.AvailableDay) error
	DeleteAvailableDay(day *models.AvailableDay) error
	GetAvailableDaysByCounselorID(counselorID uint) ([]models.AvailableDay, error)
}

type availableDayService struct {
	availableDayRepo repository.AvailableDayRepository
}

func NewAvailableDayService(availableDayRepo repository.AvailableDayRepository) AvailableDayService {
	return &availableDayService{
		availableDayRepo: availableDayRepo,
	}
}

func (s *availableDayService) CreateAvailableDay(day *models.AvailableDay) error {
	if day.ValidTo.Before(day.ValidFrom) {
		return errors.New("invalid date range")
	}
	if len(day.Weekdays) == 0 {
		return errors.New("at least one weekday must be specified")
	}
	if len(day.AvailableTimes) == 0 {
		return errors.New("at least one available time must be specified")
	}
	return s.availableDayRepo.Create(day)
}

func (s *availableDayService) CreateTmpAvailableDay(day *models.AvailableDay) error {
	if day.ValidTo.Before(day.ValidFrom) {
		return errors.New("invalid date range")
	}
	day.IsTemp = true
	day.Weekdays = []time.Weekday{}
	return s.availableDayRepo.Create(day)
}

func (s *availableDayService) UpdateTmpAvailableDay(day *models.AvailableDay) error {
	if day.ValidTo.Before(day.ValidFrom) {
		return errors.New("invalid date range")
	}
	day.IsTemp = true
	day.Weekdays = []time.Weekday{}
	return s.availableDayRepo.Update(day)
}

func (s *availableDayService) CreateHoliday(day *models.AvailableDay) error {
	if day.ValidTo.Before(day.ValidFrom) {
		return errors.New("invalid date range")
	}
	day.IsHoliday = true
	day.Weekdays = []time.Weekday{}
	day.AvailableTimes = []models.AvailableTime{}
	return s.availableDayRepo.Create(day)
}

func (s *availableDayService) UpdateHoliday(day *models.AvailableDay) error {
	if day.ValidTo.Before(day.ValidFrom) {
		return errors.New("invalid date range")
	}
	day.IsHoliday = true
	day.Weekdays = []time.Weekday{}
	day.AvailableTimes = []models.AvailableTime{}
	return s.availableDayRepo.Update(day)
}

func (s *availableDayService) UpdateAvailableDay(day *models.AvailableDay) error {
	if day.ValidTo.Before(day.ValidFrom) {
		return errors.New("invalid date range")
	}
	if len(day.Weekdays) == 0 {
		return errors.New("at least one weekday must be specified")
	}
	if len(day.AvailableTimes) == 0 {
		return errors.New("at least one available time must be specified")
	}
	return s.availableDayRepo.Update(day)
}

func (s *availableDayService) DeleteAvailableDay(day *models.AvailableDay) error {
	return s.availableDayRepo.Delete(day)
}

func (s *availableDayService) GetAvailableDaysByCounselorID(counselorID uint) ([]models.AvailableDay, error) {
	return s.availableDayRepo.GetAllByCounselorID(counselorID)
}
