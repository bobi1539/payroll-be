package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func BodyParser[T any](ctx *fiber.Ctx) *T {
	request := new(T)
	err := ctx.BodyParser(request)
	PanicIfError(err)
	return request
}

func GetParamId(ctx *fiber.Ctx, paramName string) int64 {
	id := ctx.Params(paramName)

	idInt, err := strconv.ParseInt(id, 10, 64)
	PanicIfError(err)

	return idInt
}
