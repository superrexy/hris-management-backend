package model

import (
	"hris-management/utils/model"
	"time"
)

type WorkSchedule struct {
	model.BaseModel
	ScheduleName      string     `gorm:"not null"`
	StartDate         *time.Time `gorm:"null"`
	EndDate           *time.Time `gorm:"null"`
	IsActive          bool       `gorm:"default:true"`
	WorkDays          []WorkDay
	UserWorkSchedules []UserWorkSchedule
}

type WorkDay struct {
	model.BaseModel
	WorkScheduleID uint      `gorm:"not null"`
	DayOfWeek      string    `gorm:"not null;comment:MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY,SUNDAY"`
	StartTime      time.Time `gorm:"not null"`
	EndTime        time.Time `gorm:"not null"`
	IsWorkingDay   bool      `gorm:"default:true"`
}

type UserWorkSchedule struct {
	model.BaseModel
	UserID         uint
	WorkScheduleID uint
	StartDate      *time.Time `gorm:"null"`
	EndDate        *time.Time `gorm:"null"`
}
