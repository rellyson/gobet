package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealtcheckResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func HealtcheckMiddleware(c *fiber.Ctx) error {
	c.Set("Content-type", "application/json")

	return c.Status(http.StatusOK).JSON(HealtcheckResponse{
		Message:    "Ok",
		StatusCode: http.StatusOK,
	})
}
