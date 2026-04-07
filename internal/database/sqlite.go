package database

import (
	"log"
	
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("url_shortener.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected")
	log.Println("Running Migrations")
	err = db.AutoMigrate(&models.URL{})
	if err != nil {
		log.Fatal("Failed to run migrations. \n", err)
	}

	DB = db
}
