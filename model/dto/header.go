package dto

import (
	"payroll/constant"

	"github.com/gofiber/fiber/v2"
)

type Header struct {
	UserId int64
}

func GetHeader(ctx *fiber.Ctx) Header {
	headerFromCtx := ctx.Locals(constant.HEADER)
	return headerFromCtx.(Header)
}
