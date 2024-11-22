package dto

import "time"

type StoreWorkScheduleRequest struct {
	ScheduleName string    `json:"schedule_name" validate:"required"`
	StartDate    time.Time `json:"start_date" validate:"omitempty"`
	EndDate      time.Time `json:"end_date" validate:"omitempty"`
	IsActive     bool      `json:"is_active" validate:"omitempty"`
}

type UpdateWorkScheduleRequest struct {
	ScheduleName string    `json:"schedule_name" validate:"omitempty"`
	StartDate    time.Time `json:"start_date" validate:"omitempty"`
	EndDate      time.Time `json:"end_date" validate:"omitempty"`
	IsActive     bool      `json:"is_active" validate:"omitempty"`
}

type StoreWorkDayRequest struct {
	WorkScheduleID uint      `json:"work_schedule_id" validate:"required"`
	DayOfWeek      string    `json:"day_of_week" validate:"required;oneof=MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY"`
	StartTime      time.Time `json:"start_time" validate:"required"`
	EndTime        time.Time `json:"end_time" validate:"required"`
	IsWorkingDay   bool      `json:"is_working_day" validate:"omitempty;default:true"`
}

type UpdateWorkDayRequest struct {
	WorkScheduleID uint      `json:"work_schedule_id" validate:"omitempty"`
	DayOfWeek      string    `json:"day_of_week" validate:"omitempty;oneof=MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY"`
	StartTime      time.Time `json:"start_time" validate:"omitempty"`
	EndTime        time.Time `json:"end_time" validate:"omitempty"`
	IsWorkingDay   bool      `json:"is_working_day" validate:"omitempty"`
}

type StoreUserWorkScheduleRequest struct {
	UserID         uint      `json:"user_id" validate:"required"`
	WorkScheduleID uint      `json:"work_schedule_id" validate:"required"`
	StartDate      time.Time `json:"start_date" validate:"omitempty"`
	EndDate        time.Time `json:"end_date" validate:"omitempty"`
}

type UpdateUserWorkScheduleRequest struct {
	UserID         uint      `json:"user_id" validate:"omitempty"`
	WorkScheduleID uint      `json:"work_schedule_id" validate:"omitempty"`
	StartDate      time.Time `json:"start_date" validate:"omitempty"`
	EndDate        time.Time `json:"end_date" validate:"omitempty"`
}
