package exception

import (
	"payroll/constant"
	"payroll/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = helper.GetLogger()

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Error(constant.ERROR, err)

	code := fiber.StatusInternalServerError
	message := constant.INTERNAL_SERVER_ERROR

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = err.Error()
	}

	if e, ok := err.(ErrorBusiness); ok {
		code = e.Code
		message = err.Error()
	}

	return ctx.Status(code).JSON(helper.BuildErrorResponse(code, message))
}
