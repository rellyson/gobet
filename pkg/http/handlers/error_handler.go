package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rellyson/gobet/pkg/http/errors"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	errJson := errors.ErrHttpBase{}

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	ctx.Set("Content-type", "application/json")
	errJson.Error = err.Error()
	errJson.Reason = http.StatusText(code)
	errJson.StatusCode = code

	return ctx.Status(code).JSON(errJson)
}
