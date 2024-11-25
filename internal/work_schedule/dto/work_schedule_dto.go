package dto

import "time"

type StoreWorkScheduleRequest struct {
	ScheduleName string     `json:"schedule_name" validate:"required"`
	StartDate    *time.Time `json:"start_date" validate:"omitempty,datetime"`
	EndDate      *time.Time `json:"end_date" validate:"omitempty,datetime"`
	IsActive     bool       `json:"is_active" validate:"omitempty"`
}

type UpdateWorkScheduleRequest struct {
	ScheduleName string     `json:"schedule_name" validate:"omitempty"`
	StartDate    *time.Time `json:"start_date" validate:"omitempty,datetime"`
	EndDate      *time.Time `json:"end_date" validate:"omitempty,datetime"`
	IsActive     bool       `json:"is_active" validate:"omitempty"`
}

type StoreWorkDayRequest struct {
	WorkScheduleID uint   `json:"work_schedule_id" validate:"required,number"`
	DayOfWeek      string `json:"day_of_week" validate:"required,oneof=MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY"`
	StartTime      string `json:"start_time" validate:"required,timeformat"`
	EndTime        string `json:"end_time" validate:"required,timeformat"`
	IsWorkingDay   bool   `json:"is_working_day" validate:"omitempty"`
}

type UpdateWorkDayRequest struct {
	WorkScheduleID uint   `json:"work_schedule_id" validate:"omitempty,number"`
	DayOfWeek      string `json:"day_of_week" validate:"omitempty,oneof=MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY"`
	StartTime      string `json:"start_time" validate:"omitempty,timeformat"`
	EndTime        string `json:"end_time" validate:"omitempty,timeformat"`
	IsWorkingDay   bool   `json:"is_working_day" validate:"omitempty"`
}

type StoreUserWorkScheduleRequest struct {
	UserID         uint    `json:"user_id" validate:"required,number"`
	WorkScheduleID uint    `json:"work_schedule_id" validate:"required,number"`
	StartDate      *string `json:"start_date" validate:"omitempty,datetime"`
	EndDate        *string `json:"end_date" validate:"omitempty,datetime"`
}

type UpdateUserWorkScheduleRequest struct {
	UserID         uint    `json:"user_id" validate:"omitempty,number"`
	WorkScheduleID uint    `json:"work_schedule_id" validate:"omitempty,number"`
	StartDate      *string `json:"start_date" validate:"omitempty,datetime"`
	EndDate        *string `json:"end_date" validate:"omitempty,datetime"`
}
