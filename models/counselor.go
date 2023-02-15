package models

import (
	"time"

	"gorm.io/gorm"
)

type Counselor struct {
	gorm.Model
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	Email      string `gorm:"not null;unique"`
	Password   string `gorm:"not null"`
	Profile    string
	ProfilePic string

	Phone    string `gorm:"not null;unique"`
	Address1 string `gorm:"not null"`
	Address2 string `gorm:"not null"`

	Specialty  string  `gorm:"not null"`
	Experience int     `gorm:"not null"`
	Rate       float64 `gorm:"not null"`

	Appointments    []Appointment
	AvailableDays   []AvailableDay
	CounselingMenus []CounselingMenu

	StripeAccountID string
}

func (counselor *Counselor) IsAvailableAt(day time.Time) bool {
	for _, availableDay := range counselor.AvailableDays {
		if availableDay.IsAvailableAt(day) {
			return true
		}
	}
	return false
}
