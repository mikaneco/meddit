package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	ID               int       `gorm:"primaryKey"`
	StartAt          time.Time `gorm:"not null"`
	EndAt            time.Time `gorm:"not null"`
	PaymentID        string
	PaymentState     PaymentState
	Canceled         bool
	Notes            string
	UserID           uint `gorm:"not null"`
	User             User
	CounselorID      uint `gorm:"not null"`
	Counselor        Counselor
	CounselingMenuID uint `gorm:"not null"`
	CounselingMenu   CounselingMenu
}

type PaymentState string

const (
	PaymentStateCreated   PaymentState = "created"
	PaymentStateSucceeded PaymentState = "succeeded"
	PaymentStateCanceled  PaymentState = "canceled"
)
