package config

import (
	"fmt"
	attendanceModel "hris-management/internal/attendance/model"
	userModel "hris-management/internal/user/model"
	workScheduleModel "hris-management/internal/work_schedule/model"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TAG = "Config::> "

func InitENV() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file\n", err)
	}

	log.Info(TAG, "Environment variables loaded")
}

var DB *gorm.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// MYSQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Error connecting to database\n", err)
	}

	db.AutoMigrate(&userModel.User{},
		&workScheduleModel.WorkSchedule{},
		&workScheduleModel.WorkDay{},
		&workScheduleModel.UserWorkSchedule{},
		&attendanceModel.Attendance{},
	)

	DB = db

	log.Info(TAG, "Database connected")
}
