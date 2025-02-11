package helper

import (
	"payroll/constant"
	"payroll/model/response"

	"github.com/gofiber/fiber/v3"
)

func BuildSuccessResponse(data any) response.WebResponse {
	return response.WebResponse{
		Code:    fiber.StatusOK,
		Message: constant.SUCCESS,
		Data:    data,
	}
}

func BuildErrorResponse(code int, message string) response.WebResponse {
	return response.WebResponse{
		Code:    code,
		Message: message,
	}
}
