package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
)

type UserService interface {
	Create(request *request.UserCreateRequest, header dto.Header) response.UserResponse
	Update(id int64, request *request.UserUpdateRequest, header dto.Header) response.UserResponse
	FindById(id int64) response.UserResponse
	FindByIdDomain(id int64) *domain.User
	FindAll(search *dto.Search) []response.UserResponse
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.UserResponse
}
