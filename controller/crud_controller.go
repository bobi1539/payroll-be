package controller

import "github.com/gofiber/fiber/v2"

type CrudController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}
