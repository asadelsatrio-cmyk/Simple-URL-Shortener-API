package services

import (
	"errors"

	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/models"
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/repositories"
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/utils"
)

func ShortenURL(longUrl string) (models.URL, error) {
	if longUrl == "" {
		return models.URL{}, errors.New("url cannot be empty")
	}

	shortCode := utils.GenerateShortCode(6)

	url := models.URL{
		LongURL:   longUrl,
		ShortCode: shortCode,
	}

	err := repositories.CreateURL(&url)
	if err != nil {
		return models.URL{}, err
	}

	return url, nil
}

func ResolveURL(shortCode string) (string, error) {
	url, err := repositories.GetURLByShortCode(shortCode)
	if err != nil {
		return "", errors.New("url not found")
	}

	// Increment clicks asynchronously or synchronously
	go repositories.IncrementClicks(shortCode)

	return url.LongURL, nil
}

func GetURLStats(shortCode string) (models.URL, error) {
	url, err := repositories.GetURLByShortCode(shortCode)
	if err != nil {
		return models.URL{}, errors.New("url not found")
	}
	return url, nil
}
