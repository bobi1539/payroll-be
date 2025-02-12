package main

import (
	"payroll/app"
	"payroll/constant"
	"payroll/endpoint"
	"payroll/exception"
	"payroll/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fiberApp := getFiberApp()
	fiberApp.Use(recover.New())

	db := app.NewDB()
	validate := validator.New()

	endpoint.SetRoleEndpoint(fiberApp, db, validate)

	err := fiberApp.Listen(constant.APP_HOST)
	helper.PanicIfError(err)
}

func getFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
}
