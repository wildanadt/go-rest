package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBPassword := os.Getenv("DB_PASSWORD")

	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DBHost, DBUser, DBPassword, DBName, DBPort)

	DB, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to Database")
		log.Fatal("connection error", err)
	} else {
		fmt.Println("Success connect to Database")
	}

	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.User{})

	return DB

}

func GetDBInstance() *gorm.DB {
	return DB
}
