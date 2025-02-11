package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func BodyParser[T any](ctx fiber.Ctx) *T {
	request := new(T)
	err := ctx.Bind().Body(request)
	PanicIfError(err)
	return request
}

func GetParamId(ctx fiber.Ctx, paramName string) int64 {
	id := ctx.Params(paramName)

	idInt, err := strconv.ParseInt(id, 10, 64)
	PanicIfError(err)

	return idInt
}
