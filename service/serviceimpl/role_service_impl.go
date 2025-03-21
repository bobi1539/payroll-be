package serviceimpl

import (
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/repository"
	"payroll/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewRoleServiceImpl(
	roleRepository repository.RoleRepository,
	userRepository repository.UserRepository,
	validate *validator.Validate,
) service.RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (roleService *RoleServiceImpl) Create(request *request.RoleRequest, header dto.Header) response.RoleResponse {
	roleService.validateRequest(request)
	user := roleService.findUserById(header.UserId)

	role := &domain.Role{}
	role.Name = request.Name
	helper.SetCreated(&role.BaseDomain, user)
	helper.SetUpdated(&role.BaseDomain, user)

	return response.ToRoleResponse(roleService.RoleRepository.Create(role))
}

func (roleService *RoleServiceImpl) Update(id int64, request *request.RoleRequest, header dto.Header) response.RoleResponse {
	roleService.validateRequest(request)
	user := roleService.findUserById(header.UserId)

	role := roleService.FindByIdDomain(id)
	role.Name = request.Name
	helper.SetUpdated(&role.BaseDomain, user)

	return response.ToRoleResponse(roleService.RoleRepository.Update(role))
}

func (roleService *RoleServiceImpl) FindById(id int64) response.RoleResponse {
	role := roleService.FindByIdDomain(id)
	return response.ToRoleResponse(role)
}

func (roleService *RoleServiceImpl) FindByIdDomain(id int64) *domain.Role {
	role, err := roleService.RoleRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return role
}

func (roleService *RoleServiceImpl) FindAll(search *dto.Search) []response.RoleResponse {
	roles := roleService.RoleRepository.FindAll(search)
	return response.ToRoleResponses(roles)
}

func (roleService *RoleServiceImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse {
	roles := roleService.RoleRepository.FindAllPagination(search, pagination)
	totalItem := roleService.RoleRepository.FindTotalItem(search)

	responses := response.ToRoleResponses(roles)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (roleService *RoleServiceImpl) Delete(id int64) response.RoleResponse {
	role := roleService.FindByIdDomain(id)
	roleService.RoleRepository.Delete(id)
	return response.ToRoleResponse(role)
}

func (roleService *RoleServiceImpl) validateRequest(roleRequest *request.RoleRequest) {
	err := roleService.Validate.Struct(roleRequest)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
}

func (roleService *RoleServiceImpl) findUserById(id int64) *domain.User {
	user, err := roleService.UserRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return user
}
