package workschedule

import (
	"hris-management/config"
	"hris-management/internal/work_schedule/dto"
	"hris-management/internal/work_schedule/model"
)

type WorkScheduleRepository interface {
	GetAllWorkSchedule() ([]model.WorkSchedule, error)
	GetWorkScheduleByID(id uint) (model.WorkSchedule, error)
	CreateWorkSchedule(workSchedule model.WorkSchedule) (model.WorkSchedule, error)
	UpdateWorkSchedule(workSchedule model.WorkSchedule) (model.WorkSchedule, error)
	DeleteWorkSchedule(id uint) error
}

type workScheduleRepository struct{}

// CreateWorkSchedule implements WorkScheduleRepository.
func (w *workScheduleRepository) CreateWorkSchedule(workSchedule model.WorkSchedule) (model.WorkSchedule, error) {
	result := config.DB.Create(&workSchedule)
	if result.Error != nil {
		return model.WorkSchedule{}, result.Error
	}

	return workSchedule, nil
}

// DeleteWorkSchedule implements WorkScheduleRepository.
func (w *workScheduleRepository) DeleteWorkSchedule(id uint) error {
	var workSchedule model.WorkSchedule

	result := config.DB.Where("id = ?", id).Delete(&workSchedule)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// GetAllWorkSchedule implements WorkScheduleRepository.
func (w *workScheduleRepository) GetAllWorkSchedule() ([]model.WorkSchedule, error) {
	var workSchedules []model.WorkSchedule
	result := config.DB.Find(&workSchedules)
	if result.Error != nil {
		return nil, result.Error
	}

	return workSchedules, nil
}

// GetWorkScheduleByID implements WorkScheduleRepository.
func (w *workScheduleRepository) GetWorkScheduleByID(id uint) (model.WorkSchedule, error) {
	var workSchedule model.WorkSchedule

	result := config.DB.Where("id = ?", id).First(&workSchedule)
	if result.Error != nil {
		return model.WorkSchedule{}, result.Error
	}

	return workSchedule, nil
}

// UpdateWorkSchedule implements WorkScheduleRepository.
func (w *workScheduleRepository) UpdateWorkSchedule(workSchedule model.WorkSchedule) (model.WorkSchedule, error) {
	data := dto.UpdateWorkScheduleRequest{
		ScheduleName: workSchedule.ScheduleName,
		StartDate:    workSchedule.StartDate,
		EndDate:      workSchedule.EndDate,
		IsActive:     workSchedule.IsActive,
	}

	result := config.DB.Model(&workSchedule).Updates(data)
	if result.Error != nil {
		return model.WorkSchedule{}, result.Error
	}

	return workSchedule, nil
}

func NewWorkScheduleRepository() WorkScheduleRepository {
	return &workScheduleRepository{}
}

type WorkDayRepository interface {
	CreateWorkDay(workDay model.WorkDay) (model.WorkDay, error)
	DeleteWorkDay(id uint) error
	GetAllWorkDay() ([]model.WorkDay, error)
	GetWorkDayByID(id uint) (model.WorkDay, error)
	UpdateWorkDay(workDay model.WorkDay) (model.WorkDay, error)
}

type workDayRepository struct{}

// CreateWorkDay implements WorkDayRepository.
func (w *workDayRepository) CreateWorkDay(workDay model.WorkDay) (model.WorkDay, error) {
	result := config.DB.Create(&workDay)
	if result.Error != nil {
		return model.WorkDay{}, result.Error
	}

	return workDay, nil
}

// DeleteWorkDay implements WorkDayRepository.
func (w *workDayRepository) DeleteWorkDay(id uint) error {
	var workDay model.WorkDay

	result := config.DB.Where("id = ?", id).Delete(&workDay)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllWorkDay implements WorkDayRepository.
func (w *workDayRepository) GetAllWorkDay() ([]model.WorkDay, error) {
	var workDays []model.WorkDay
	result := config.DB.Find(&workDays)
	if result.Error != nil {
		return nil, result.Error
	}

	return workDays, nil
}

// GetWorkDayByID implements WorkDayRepository.
func (w *workDayRepository) GetWorkDayByID(id uint) (model.WorkDay, error) {
	var workDay model.WorkDay

	result := config.DB.Where("id = ?", id).First(&workDay)
	if result.Error != nil {
		return model.WorkDay{}, result.Error
	}

	return workDay, nil
}

// UpdateWorkDay implements WorkDayRepository.
func (w *workDayRepository) UpdateWorkDay(workDay model.WorkDay) (model.WorkDay, error) {
	data := dto.UpdateWorkDayRequest{
		WorkScheduleID: workDay.WorkScheduleID,
		DayOfWeek:      workDay.DayOfWeek,
		StartTime:      workDay.StartTime,
		EndTime:        workDay.EndTime,
		IsWorkingDay:   workDay.IsWorkingDay,
	}

	result := config.DB.Model(&workDay).Updates(data)
	if result.Error != nil {
		return model.WorkDay{}, result.Error
	}

	return workDay, nil
}

func NewWorkDayRepository() WorkDayRepository {
	return &workDayRepository{}
}

type UserWorkScheduleRepository interface {
	CreateUserWorkSchedule(userWorkSchedule model.UserWorkSchedule) (model.UserWorkSchedule, error)
	DeleteUserWorkSchedule(id uint) error
	GetAllUserWorkSchedule() ([]model.UserWorkSchedule, error)
	GetUserWorkScheduleByID(id uint) (model.UserWorkSchedule, error)
	UpdateUserWorkSchedule(userWorkSchedule model.UserWorkSchedule) (model.UserWorkSchedule, error)
}

type userWorkScheduleRepository struct{}

// CreateUserWorkSchedule implements UserWorkScheduleRepository.
func (u *userWorkScheduleRepository) CreateUserWorkSchedule(userWorkSchedule model.UserWorkSchedule) (model.UserWorkSchedule, error) {
	result := config.DB.Create(&userWorkSchedule)
	if result.Error != nil {
		return model.UserWorkSchedule{}, result.Error
	}

	return userWorkSchedule, nil
}

// DeleteUserWorkSchedule implements UserWorkScheduleRepository.
func (u *userWorkScheduleRepository) DeleteUserWorkSchedule(id uint) error {
	var userWorkSchedule model.UserWorkSchedule

	result := config.DB.Where("id = ?", id).Delete(&userWorkSchedule)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllUserWorkSchedule implements UserWorkScheduleRepository.
func (u *userWorkScheduleRepository) GetAllUserWorkSchedule() ([]model.UserWorkSchedule, error) {
	var userWorkSchedules []model.UserWorkSchedule
	result := config.DB.Find(&userWorkSchedules)
	if result.Error != nil {
		return nil, result.Error
	}

	return userWorkSchedules, nil
}

// GetUserWorkScheduleByID implements UserWorkScheduleRepository.
func (u *userWorkScheduleRepository) GetUserWorkScheduleByID(id uint) (model.UserWorkSchedule, error) {
	var userWorkSchedule model.UserWorkSchedule

	result := config.DB.Where("id = ?", id).First(&userWorkSchedule)
	if result.Error != nil {
		return model.UserWorkSchedule{}, result.Error
	}

	return userWorkSchedule, nil
}

// UpdateUserWorkSchedule implements UserWorkScheduleRepository.
func (u *userWorkScheduleRepository) UpdateUserWorkSchedule(userWorkSchedule model.UserWorkSchedule) (model.UserWorkSchedule, error) {
	result := config.DB.Model(&userWorkSchedule).Updates(&userWorkSchedule)
	if result.Error != nil {
		return model.UserWorkSchedule{}, result.Error
	}

	return userWorkSchedule, nil
}

func NewUserWorkScheduleRepository() UserWorkScheduleRepository {
	return &userWorkScheduleRepository{}
}
