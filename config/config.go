package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database\n", err)
	}

	DB = db

	log.Info(TAG, "Database connected")
}
