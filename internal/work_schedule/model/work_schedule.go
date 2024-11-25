package model

import (
	"hris-management/utils/model"
	"time"
)

type WorkSchedule struct {
	model.BaseModel
	ScheduleName      string             `gorm:"not null" json:"schedule_name"`
	StartDate         *time.Time         `gorm:"null" json:"start_date"`
	EndDate           *time.Time         `gorm:"null" json:"end_date"`
	IsActive          bool               `gorm:"default:true" json:"is_active"`
	WorkDays          []WorkDay          `json:"work_days" gorm:"constraint:OnDelete:CASCADE"`
	UserWorkSchedules []UserWorkSchedule `json:"user_work_schedules" gorm:"constraint:OnDelete:CASCADE"`
}

type WorkDay struct {
	model.BaseModel
	WorkScheduleID uint   `gorm:"not null" json:"work_schedule_id"`
	DayOfWeek      string `gorm:"not null;comment:MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY,SUNDAY" json:"day_of_week"`
	StartTime      string `gorm:"not null" json:"start_time"`
	EndTime        string `gorm:"not null" json:"end_time"`
	IsWorkingDay   bool   `gorm:"default:true" json:"is_working_day"`
}

type UserWorkSchedule struct {
	model.BaseModel
	UserID         uint       `gorm:"not null" json:"user_id"`
	WorkScheduleID uint       `gorm:"not null" json:"work_schedule_id"`
	StartDate      *time.Time `gorm:"null" json:"start_date"`
	EndDate        *time.Time `gorm:"null" json:"end_date"`
}
