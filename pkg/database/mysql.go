package database

import (
	"car-rental/internal/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToMySQL() *gorm.DB{
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUSER := os.Getenv("DB_USER")
	dbPASS := os.Getenv("DB_PASSWORD")
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUSER, dbPASS, dbHOST, dbPORT, dbNAME)

	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return nil
	}
	db.AutoMigrate(&models.Users{}, &models.Cars{}, &models.Customers{}, &models.Bookings{})
	DB = db
	return db
}