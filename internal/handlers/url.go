package handlers

import (
	"github.com/asadelsatrio/Simple-URL-Shortener-API/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

func ShortenURL(c *fiber.Ctx) error {
	var req ShortenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON body",
		})
	}

	url, err := services.ShortenURL(req.LongURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(url)
}

func RedirectURL(c *fiber.Ctx) error {
	shortCode := c.Params("code")

	longUrl, err := services.ResolveURL(shortCode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect(longUrl, fiber.StatusMovedPermanently)
}

func GetURLStats(c *fiber.Ctx) error {
	shortCode := c.Params("code")

	url, err := services.GetURLStats(shortCode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(url)
}
