package database

import (
	"fmt"
	"log"
	"os"
	"payment-service/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")

  dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
  db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }
  
	db.AutoMigrate(&model.User{}, &model.Payment{})
  fmt.Println("Successfully connected to database!")
}

func GetDB() *gorm.DB{
    return db
}