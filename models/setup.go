package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	errDb error
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// fmt.Println("Error env: ", err)
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("dbUser")
	dbPassword := os.Getenv("dbPassword")
	dbName := os.Getenv("dbName")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", dbUser, dbPassword, dbName)
	db, errDb = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDb != nil {
		panic("Failed to connect to database!")
	}

	db.Debug().AutoMigrate(
		&Todo{},
	)
}

func GetDB() *gorm.DB {
	return db
}
