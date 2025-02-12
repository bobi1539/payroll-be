package exception

import (
	"payroll/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = helper.GetLogger()

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Error("Error : ", err)

	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if e, ok := err.(ErrorBusiness); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(helper.BuildErrorResponse(code, err.Error()))
}
