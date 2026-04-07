package repositories

import (
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/database"
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/models"
)

func CreateURL(url *models.URL) error {
	return database.DB.Create(url).Error
}

func GetURLByShortCode(shortCode string) (models.URL, error) {
	var url models.URL
	err := database.DB.Where("short_code = ?", shortCode).First(&url).Error
	return url, err
}

func IncrementClicks(shortCode string) error {
	return database.DB.Model(&models.URL{}).Where("short_code = ?", shortCode).UpdateColumn("clicks", database.DB.Raw("clicks + ?", 1)).Error
}
