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
	"gorm.io/gorm"

	_ "payroll/docs"
)

// @title		PAYROLL API
// @version		1.0
// @description	This is a payroll api service.
// @BasePath	/api
func main() {
	fiberApp := getFiberApp()
	db := app.NewDB()
	validate := validator.New()

	setEndpoint(fiberApp, db, validate)

	err := fiberApp.Listen(constant.APP_HOST)
	helper.PanicIfError(err)
}

func getFiberApp() *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	fiberApp.Use(recover.New())
	return fiberApp
}

func setEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	endpoint.SetSwaggerEndpoint(fiberApp)
	endpoint.SetRoleEndpoint(fiberApp, db, validate)
	endpoint.SetUserEndpoint(fiberApp, db, validate)
	endpoint.SetAuthEndpoint(fiberApp, db, validate)
}
