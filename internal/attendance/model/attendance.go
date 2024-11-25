package model

import (
	"hris-management/utils/model"
	"time"
)

type Attendance struct {
	model.BaseModel
	UserID            uint       `json:"user_id" gorm:"not null"`
	CheckIn           time.Time  `json:"check_in" gorm:"not null"`
	CheckOut          *time.Time `json:"check_out" gorm:"default:null"`
	CheckInLatitude   float64    `json:"check_in_latitude" gorm:"not null"`
	CheckInLongitude  float64    `json:"check_in_longitude" gorm:"not null"`
	CheckOutLatitude  *float64   `json:"check_out_latitude" gorm:"default:null"`
	CheckOutLongitude *float64   `json:"check_out_longitude" gorm:"default:null"`
	Status            *string    `json:"status" gorm:"default:null"`
}
