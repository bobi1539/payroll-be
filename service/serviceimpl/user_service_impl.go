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

type UserServiceImpl struct {
	UserRepository        repository.UserRepository
	Validate              *validator.Validate
	RoleService           service.RoleService
	UserValidationService service.UserValidationService
}

func NewUserService(
	userRepository repository.UserRepository,
	validate *validator.Validate,
	roleService service.RoleService,
	userValidationService service.UserValidationService,
) service.UserService {
	return &UserServiceImpl{
		UserRepository:        userRepository,
		Validate:              validate,
		RoleService:           roleService,
		UserValidationService: userValidationService,
	}
}

func (userService *UserServiceImpl) Create(request *request.UserCreateRequest) response.UserResponse {
	userService.validateCreateUser(request)

	user := &domain.User{}
	user.Name = request.Name
	user.Username = request.Username
	user.Role = userService.findRoleById(request.RoleId)
	helper.SetCreated(&user.BaseDomain)
	helper.SetUpdated(&user.BaseDomain)

	return response.ToUserResponse(userService.UserRepository.Create(user))
}

func (userService *UserServiceImpl) Update(id int64, request *request.UserUpdateRequest) response.UserResponse {
	err := userService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)

	user := userService.FindByIdDomain(id)
	userService.UserValidationService.ValidateUpdateUsername(request.Username, user)
	user.Name = request.Name
	user.Username = request.Username
	user.Role = userService.findRoleById(request.RoleId)
	helper.SetUpdated(&user.BaseDomain)

	return response.ToUserResponse(userService.UserRepository.Update(user))
}

func (userService *UserServiceImpl) FindById(id int64) response.UserResponse {
	user := userService.FindByIdDomain(id)
	return response.ToUserResponse(user)
}

func (userService *UserServiceImpl) FindByIdDomain(id int64) *domain.User {
	user, err := userService.UserRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return user
}

func (userService *UserServiceImpl) FindAll(search *dto.Search) []response.UserResponse {
	users := userService.UserRepository.FindAll(search)
	return response.ToUserResponses(users)
}

func (userService *UserServiceImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse {
	users := userService.UserRepository.FindAllPagination(search, pagination)
	totalItem := userService.UserRepository.FindTotalItem()

	responses := response.ToUserResponses(users)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (userService *UserServiceImpl) Delete(id int64) response.UserResponse {
	user := userService.FindByIdDomain(id)
	userService.UserRepository.Delete(id)
	return response.ToUserResponse(user)
}

func (userService *UserServiceImpl) findRoleById(id int64) *domain.Role {
	return userService.RoleService.FindByIdDomain(id)
}

func (userService *UserServiceImpl) validateCreateUser(request *request.UserCreateRequest) {
	err := userService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	userService.UserValidationService.ValidateCreateUsername(request.Username)
	userService.UserValidationService.ValidatePassword(request.Password)
}
