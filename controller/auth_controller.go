package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	LoginRefreshToken(ctx *fiber.Ctx) error
}
