package workschedule

import (
	"fmt"
	user "hris-management/internal/user"
	"hris-management/internal/work_schedule/dto"
	"hris-management/internal/work_schedule/model"
	"hris-management/utils/exception"
	"time"

	"github.com/gofiber/fiber/v2"
)

type WorkScheduleService interface {
	GetAllWorkSchedule() ([]model.WorkSchedule, error)
	GetWorkScheduleByID(id uint) (model.WorkSchedule, error)
	CreateWorkSchedule(workSchedule dto.StoreWorkScheduleRequest) (model.WorkSchedule, error)
	UpdateWorkSchedule(workSchedule dto.UpdateWorkScheduleRequest, id uint) (model.WorkSchedule, error)
	DeleteWorkSchedule(id uint) error
}

type workScheduleService struct {
	workScheduleRepository WorkScheduleRepository
}

// CreateWorkSchedule implements WorkScheduleService.
func (w *workScheduleService) CreateWorkSchedule(workSchedule dto.StoreWorkScheduleRequest) (model.WorkSchedule, error) {
	payload := model.WorkSchedule{
		ScheduleName: workSchedule.ScheduleName,
		StartDate:    workSchedule.StartDate,
		EndDate:      workSchedule.EndDate,
		IsActive:     workSchedule.IsActive,
	}

	if payload.StartDate != nil && payload.EndDate != nil {
		if payload.StartDate.After(*payload.EndDate) {
			return model.WorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "End date must be greater than start date", nil)
		} else if payload.StartDate.Equal(*payload.EndDate) {
			return model.WorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Start date and end date must not be the same", nil)
		}
	}

	workScheduleData, err := w.workScheduleRepository.CreateWorkSchedule(payload)
	if err != nil {
		return model.WorkSchedule{}, err
	}

	return workScheduleData, nil
}

// DeleteWorkSchedule implements WorkScheduleService.
func (w *workScheduleService) DeleteWorkSchedule(id uint) error {
	schedule, _ := w.GetWorkScheduleByID(id)
	if schedule.ID == 0 {
		return exception.NewServiceError(fiber.StatusNotFound, "Work Schedule not found", nil)
	}

	err := w.workScheduleRepository.DeleteWorkSchedule(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllWorkSchedule implements WorkScheduleService.
func (w *workScheduleService) GetAllWorkSchedule() ([]model.WorkSchedule, error) {
	workScheduleData, err := w.workScheduleRepository.GetAllWorkSchedule()
	if err != nil {
		return nil, err
	}

	return workScheduleData, nil
}

// GetWorkScheduleByID implements WorkScheduleService.
func (w *workScheduleService) GetWorkScheduleByID(id uint) (model.WorkSchedule, error) {
	workScheduleData, err := w.workScheduleRepository.GetWorkScheduleByID(id)
	if err != nil {
		return model.WorkSchedule{}, err
	}

	if workScheduleData.ID == 0 {
		return model.WorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "Work Schedule not found", nil)
	}

	return workScheduleData, nil
}

// UpdateWorkSchedule implements WorkScheduleService.
func (w *workScheduleService) UpdateWorkSchedule(workSchedule dto.UpdateWorkScheduleRequest, id uint) (model.WorkSchedule, error) {
	schedule, _ := w.GetWorkScheduleByID(id)
	if schedule.ID == 0 {
		return model.WorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "Work Schedule not found", nil)
	}

	if workSchedule.StartDate != nil && workSchedule.EndDate != nil {
		schedule.StartDate = workSchedule.StartDate
		schedule.EndDate = workSchedule.EndDate

		if workSchedule.StartDate.After(*workSchedule.EndDate) {
			return model.WorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "End date must be greater than start date", nil)
		} else if workSchedule.StartDate.Equal(*workSchedule.EndDate) {
			return model.WorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Start date and end date must not be the same", nil)
		}
	}

	schedule.ScheduleName = workSchedule.ScheduleName
	schedule.IsActive = workSchedule.IsActive

	workScheduleData, err := w.workScheduleRepository.UpdateWorkSchedule(schedule)
	if err != nil {
		return model.WorkSchedule{}, err
	}

	return workScheduleData, nil
}

func NewWorkScheduleService(workScheduleRepository WorkScheduleRepository) WorkScheduleService {
	return &workScheduleService{
		workScheduleRepository: workScheduleRepository,
	}
}

type WorkDayService interface {
	CreateWorkDay(workDay dto.StoreWorkDayRequest) (model.WorkDay, error)
	DeleteWorkDay(id uint) error
	GetAllWorkDay() ([]model.WorkDay, error)
	GetWorkDayByID(id uint) (model.WorkDay, error)
	UpdateWorkDay(workDay dto.UpdateWorkDayRequest, id uint) (model.WorkDay, error)
}

type workDayService struct {
	workScheduleRepository WorkScheduleRepository
	workDayRepository      WorkDayRepository
}

// CreateWorkDay implements WorkDayService.
func (w *workDayService) CreateWorkDay(workDay dto.StoreWorkDayRequest) (model.WorkDay, error) {
	workSchedule, _ := w.workScheduleRepository.GetWorkScheduleByID(workDay.WorkScheduleID)
	if workSchedule.ID == 0 {
		return model.WorkDay{}, exception.NewServiceError(fiber.StatusNotFound, "Work Schedule not found", nil)
	}

	payload := model.WorkDay{
		WorkScheduleID: workSchedule.ID,
		DayOfWeek:      workDay.DayOfWeek,
		StartTime:      workDay.StartTime,
		EndTime:        workDay.EndTime,
		IsWorkingDay:   workDay.IsWorkingDay,
	}

	workDayData, err := w.workDayRepository.CreateWorkDay(payload)
	if err != nil {
		return model.WorkDay{}, err
	}

	return workDayData, nil
}

// DeleteWorkDay implements WorkDayService.
func (w *workDayService) DeleteWorkDay(id uint) error {
	workDay, err := w.GetWorkDayByID(id)
	if err != nil {
		return err
	}

	err = w.workDayRepository.DeleteWorkDay(workDay.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllWorkDay implements WorkDayService.
func (w *workDayService) GetAllWorkDay() ([]model.WorkDay, error) {
	result, err := w.workDayRepository.GetAllWorkDay()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetWorkDayByID implements WorkDayService.
func (w *workDayService) GetWorkDayByID(id uint) (model.WorkDay, error) {
	result, err := w.workDayRepository.GetWorkDayByID(id)
	if err != nil {
		return model.WorkDay{}, err
	}

	if result.ID == 0 {
		return model.WorkDay{}, exception.NewServiceError(fiber.StatusNotFound, "Work Day not found", nil)
	}

	return result, nil
}

// UpdateWorkDay implements WorkDayService.
func (w *workDayService) UpdateWorkDay(workDay dto.UpdateWorkDayRequest, id uint) (model.WorkDay, error) {
	workDayData, err := w.GetWorkDayByID(id)
	if err != nil {
		return model.WorkDay{}, err
	}

	workDayData.WorkScheduleID = workDay.WorkScheduleID
	workDayData.DayOfWeek = workDay.DayOfWeek
	workDayData.IsWorkingDay = workDay.IsWorkingDay
	workDayData.StartTime = workDay.StartTime
	workDayData.EndTime = workDay.EndTime

	result, err := w.workDayRepository.UpdateWorkDay(workDayData)
	if err != nil {
		return model.WorkDay{}, err
	}

	return result, nil
}

func NewWorkDayService(workDayRepository WorkDayRepository, workScheduleRepository WorkScheduleRepository) WorkDayService {
	return &workDayService{
		workDayRepository:      workDayRepository,
		workScheduleRepository: workScheduleRepository,
	}
}

type UserWorkScheduleService interface {
	CreateUserWorkSchedule(userWorkSchedule dto.StoreUserWorkScheduleRequest) (model.UserWorkSchedule, error)
	DeleteUserWorkSchedule(id uint) error
	GetAllUserWorkSchedule() ([]model.UserWorkSchedule, error)
	GetUserWorkScheduleByID(id uint) (model.UserWorkSchedule, error)
	UpdateUserWorkSchedule(userWorkSchedule dto.UpdateUserWorkScheduleRequest, id uint) (model.UserWorkSchedule, error)
}

type userWorkScheduleService struct {
	userWorkScheduleRepository UserWorkScheduleRepository
	workScheduleRepository     WorkScheduleRepository
	userRepository             user.UserRepository
}

// CreateUserWorkSchedule implements UserWorkScheduleService.
func (u *userWorkScheduleService) CreateUserWorkSchedule(userWorkSchedule dto.StoreUserWorkScheduleRequest) (model.UserWorkSchedule, error) {
	payload := model.UserWorkSchedule{
		UserID:         userWorkSchedule.UserID,
		WorkScheduleID: userWorkSchedule.WorkScheduleID,
	}

	if userWorkSchedule.StartDate != nil && userWorkSchedule.EndDate != nil {
		startDate, err := time.Parse(time.RFC3339, *userWorkSchedule.StartDate)
		if err != nil {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Invalid start date", nil)
		}

		payload.StartDate = &startDate

		endDate, err := time.Parse(time.RFC3339, *userWorkSchedule.EndDate)
		if err != nil {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Invalid end date", nil)
		}

		payload.EndDate = &endDate
	}

	user, _ := u.userRepository.GetByAttribute("id", userWorkSchedule.UserID)
	if user.ID == 0 {
		return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "User not found", nil)
	}

	if user.IsActive == false {
		return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "User is not active", nil)
	}

	if payload.StartDate != nil && payload.EndDate != nil {
		if payload.StartDate.After(*payload.EndDate) {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "End date must be greater than start date", nil)
		} else if payload.StartDate.Equal(*payload.EndDate) {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Start date and end date must not be the same", nil)
		}
	}

	workSchedule, _ := u.workScheduleRepository.GetWorkScheduleByID(userWorkSchedule.WorkScheduleID)
	if workSchedule.ID == 0 {
		return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "Work Schedule not found", nil)
	}

	if workSchedule.IsActive == false {
		return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Work Schedule is not active", nil)
	}

	userWorkScheduleData, err := u.userWorkScheduleRepository.CreateUserWorkSchedule(payload)
	if err != nil {
		return model.UserWorkSchedule{}, err
	}

	return userWorkScheduleData, nil
}

// DeleteUserWorkSchedule implements UserWorkScheduleService.
func (u *userWorkScheduleService) DeleteUserWorkSchedule(id uint) error {
	userWorkScheduleData, err := u.GetUserWorkScheduleByID(id)
	if err != nil {
		return err
	}

	err = u.userWorkScheduleRepository.DeleteUserWorkSchedule(userWorkScheduleData.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUserWorkSchedule implements UserWorkScheduleService.
func (u *userWorkScheduleService) GetAllUserWorkSchedule() ([]model.UserWorkSchedule, error) {
	userWorkScheduleData, err := u.userWorkScheduleRepository.GetAllUserWorkSchedule()
	if err != nil {
		return nil, err
	}

	return userWorkScheduleData, nil
}

// GetUserWorkScheduleByID implements UserWorkScheduleService.
func (u *userWorkScheduleService) GetUserWorkScheduleByID(id uint) (model.UserWorkSchedule, error) {
	userWorkScheduleData, _ := u.userWorkScheduleRepository.GetUserWorkScheduleByID(id)

	if userWorkScheduleData.ID == 0 {
		return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "User Work Schedule not found", nil)
	}

	return userWorkScheduleData, nil
}

// UpdateUserWorkSchedule implements UserWorkScheduleService.
func (u *userWorkScheduleService) UpdateUserWorkSchedule(userWorkSchedule dto.UpdateUserWorkScheduleRequest, id uint) (model.UserWorkSchedule, error) {

	userWorkScheduleData, _ := u.GetUserWorkScheduleByID(id)
	if userWorkScheduleData.ID == 0 {
		return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "User Work Schedule not found", nil)
	}

	if userWorkSchedule.WorkScheduleID != userWorkScheduleData.WorkScheduleID {
		workSchedule, _ := u.workScheduleRepository.GetWorkScheduleByID(userWorkSchedule.WorkScheduleID)
		if workSchedule.ID == 0 {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "Work Schedule not found", nil)
		}

		if workSchedule.IsActive == false {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Work Schedule is not active", nil)
		}

		userWorkScheduleData.WorkScheduleID = userWorkSchedule.WorkScheduleID
	}

	if userWorkSchedule.UserID != userWorkScheduleData.UserID {
		user, _ := u.userRepository.GetByAttribute("id", userWorkSchedule.UserID)
		if user.ID == 0 {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusNotFound, "User not found", nil)
		}

		if user.IsActive == false {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "User is not active", nil)
		}

		userWorkScheduleData.UserID = userWorkSchedule.UserID
	}

	if userWorkSchedule.StartDate != nil && userWorkSchedule.EndDate != nil {
		startDate, err := time.Parse(time.RFC3339, *userWorkSchedule.StartDate)
		if err != nil {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Invalid start date", nil)
		}

		userWorkScheduleData.StartDate = &startDate

		endDate, err := time.Parse(time.RFC3339, *userWorkSchedule.EndDate)
		if err != nil {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Invalid end date", nil)
		}

		userWorkScheduleData.EndDate = &endDate
	}

	if userWorkSchedule.StartDate != nil && userWorkSchedule.EndDate != nil {
		if userWorkScheduleData.StartDate.After(*userWorkScheduleData.EndDate) {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "End date must be greater than start date", nil)
		} else if userWorkScheduleData.StartDate.Equal(*userWorkScheduleData.EndDate) {
			return model.UserWorkSchedule{}, exception.NewServiceError(fiber.StatusBadRequest, "Start date and end date must not be the same", nil)
		}
	}

	fmt.Println(userWorkScheduleData)

	result, err := u.userWorkScheduleRepository.UpdateUserWorkSchedule(userWorkScheduleData)
	if err != nil {
		return model.UserWorkSchedule{}, err
	}

	return result, nil
}

func NewUserWorkScheduleService(userWorkScheduleRepository UserWorkScheduleRepository, workScheduleRepository WorkScheduleRepository, userRepository user.UserRepository) UserWorkScheduleService {
	return &userWorkScheduleService{
		userWorkScheduleRepository: userWorkScheduleRepository,
		workScheduleRepository:     workScheduleRepository,
		userRepository:             userRepository,
	}
}
