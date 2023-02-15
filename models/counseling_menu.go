package models

import (
	"time"

	"gorm.io/gorm"
)

type CounselingMenu struct {
	gorm.Model
	Name        string	`gorm:"not null"`
	Description string	`gorm:"not null"`
	Price       int	`gorm:"not null"`
	Duration    time.Duration	`gorm:"not null"`
	CounselorID uint	`gorm:"not null"`
	Counselor   Counselor
}
