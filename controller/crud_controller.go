package controller

import "github.com/gofiber/fiber/v3"

type CrudController interface {
	Create(ctx fiber.Ctx) error
	Update(ctx fiber.Ctx) error
	FindById(ctx fiber.Ctx) error
}
