package service

import (
	"payroll/model/request"
	"payroll/model/response"
)

type AuthService interface {
	Login(request *request.LoginRequest) response.LoginResponse
}
