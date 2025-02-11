package exception

import (
	"payroll/helper"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if e, ok := err.(ErrorBusiness); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(helper.BuildErrorResponse(code, err.Error()))
}
