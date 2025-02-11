package helper

import (
	"payroll/constant"
	"payroll/model/response"
)

func BuildSuccessResponse(data any) response.WebResponse {
	return response.WebResponse{
		Code:    200,
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
