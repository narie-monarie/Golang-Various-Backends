package config

import (
	"log"
	"narie/monarie-project/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION")), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Post{})
	DB = db
}
