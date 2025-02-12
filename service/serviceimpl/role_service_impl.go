package serviceimpl

import (
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/repository"
	"payroll/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, validate *validator.Validate) service.RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		Validate:       validate,
	}
}

func (roleService *RoleServiceImpl) Create(request *request.RoleRequest) response.RoleResponse {
	roleService.validateRequest(request)

	role := &domain.Role{}
	role.Name = request.Name
	helper.SetCreated(&role.BaseDomain)
	helper.SetUpdated(&role.BaseDomain)

	role = roleService.RoleRepository.Create(role)
	return response.ToRoleResponse(role)
}

func (roleService *RoleServiceImpl) Update(id int64, request *request.RoleRequest) response.RoleResponse {
	roleService.validateRequest(request)

	role := roleService.FindByIdDomain(id)
	role.Name = request.Name
	helper.SetUpdated(&role.BaseDomain)

	role = roleService.RoleRepository.Update(role)
	return response.ToRoleResponse(role)
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

func (roleService *RoleServiceImpl) validateRequest(roleRequest *request.RoleRequest) {
	err := roleService.Validate.Struct(roleRequest)
	helper.PanicIfError(err)
}
