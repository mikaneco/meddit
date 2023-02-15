package models

import (
	"time"

	"gorm.io/gorm"
)

type AvailableDay struct {
	gorm.Model
	ID        int
	ValidFrom time.Time `gorm:"not null; COMMENT:'予約枠の適用開始日"`
	ValidTo   time.Time `gorm:"not null; COMMENT:'予約枠の適用終了日"`

	Weekdays []time.Weekday `gorm:"type:varchar(255); COMMENT:'予約枠の適用曜日'"`

	IsTemp    bool `gorm:"COMMENT:'臨時設定の場合true'"`
	IsHoliday bool `gorm:"COMMENT:'休暇の場合true'"`

	AvailableTimes []AvailableTime

	CounselorID uint `gorm:"not null"`
	Counselor   Counselor
}

type AvailableTime struct {
	gorm.Model
	ID int

	StartAt time.Time `gorm:"not null"`
	EndAt   time.Time `gorm:"not null"`

	AvailableDayID uint `gorm:"not null"`
	AvailableDay   AvailableDay
}

func (availableDay *AvailableDay) IsAvailableAt(day time.Time) bool {
	if !availableDay.IsTemp && availableDay.IsHoliday {
		return false
	}

	if availableDay.ValidFrom.After(day) || availableDay.ValidTo.Before(day) {
		return false
	}

	for _, weekday := range availableDay.Weekdays {
		if weekday == day.Weekday() {
			return true
		}
	}

	return false
}

//月曜日から金曜日まで12:00-13:00の予約枠を設定する例
/*
availableDay := models.AvailableDay{
	ValidFrom: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
	ValidTo:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.Local),
	Weekdays:  []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	IsTemp:    false,
	IsHoliday: false,
	Counselor: counselor,
	AvailableTimes: []models.AvailableTime{
		{
			StartAt: time.Date(0, 0, 0, 12, 0, 0, 0, time.Local),
			EndAt:   time.Date(0, 0, 0, 13, 0, 0, 0, time.Local),
		}
	}
}
*/

//月曜日から金曜日まで12:00-13:00, 15:00-16:00の予約枠を設定する例
/*
availableDay := models.AvailableDay{
	ValidFrom: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
	ValidTo:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.Local),
	Weekdays:  []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	IsTemp:    false,
	IsHoliday: false,
	Counselor: counselor,
	AvailableTimes: []models.AvailableTime{
		{
			StartAt: time.Date(0, 0, 0, 12, 0, 0, 0, time.Local),
			EndAt:   time.Date(0, 0, 0, 13, 0, 0, 0, time.Local),
		},
		{
			StartAt: time.Date(0, 0, 0, 15, 0, 0, 0, time.Local),
			EndAt:   time.Date(0, 0, 0, 16, 0, 0, 0, time.Local),
		}
	}
}
*/

//臨時的に土曜日の予約枠を設定する例
/*
availableDay := models.AvailableDay{
	ValidFrom: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
	ValidTo:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.Local),
	Weekdays:  []time.Weekday{time.Saturday},
	IsTemp:    true,
	IsHoliday: false,
	Counselor: counselor,
	AvailableTimes: []models.AvailableTime{
		{
			StartAt: time.Date(0, 0, 0, 12, 0, 0, 0, time.Local),
			EndAt:   time.Date(0, 0, 0, 13, 0, 0, 0, time.Local),
		}
	}
}
*/

//臨時的に12月24日から31日までの休暇を設定する例
/*
availableDay := models.AvailableDay{
	ValidFrom: time.Date(2021, 12, 24, 0, 0, 0, 0, time.Local),
	ValidTo:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.Local),
	Weekdays:  []time.Weekday{},
	IsTemp:    true,
	IsHoliday: true,
	Counselor: counselor,
	AvailableTimes: []models.AvailableTime{}
}
*/
