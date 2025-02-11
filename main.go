package main

import (
	"payroll/app"
	"payroll/endpoint"
	"payroll/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func main() {
	fiberApp := fiber.New()
	db := app.NewDB()
	validate := validator.New()

	endpoint.SetRoleEndpoint(fiberApp, db, validate)

	err := fiberApp.Listen("localhost:3000")
	helper.PanicIfError(err)
}
