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

type AllowanceTypeServiceImpl struct {
	AllowanceTypeRepository repository.AllowanceTypeRepository
	UserService             service.UserService
	Validate                *validator.Validate
}

func NewAllowanceTypeServiceImpl(
	allowanceTypeRepository repository.AllowanceTypeRepository,
	userService service.UserService,
	validate *validator.Validate,
) service.AllowanceTypeService {
	return &AllowanceTypeServiceImpl{
		AllowanceTypeRepository: allowanceTypeRepository,
		UserService:             userService,
		Validate:                validate,
	}
}

func (atService *AllowanceTypeServiceImpl) Create(request *request.AllowanceTypeRequest, header dto.Header) response.AllowanceTypeResponse {
	atService.validateRequest(request)
	user := atService.UserService.FindByIdDomain(header.UserId)

	allowanceType := &domain.AllowanceType{}
	allowanceType.Name = request.Name
	helper.SetCreated(&allowanceType.BaseDomain, user)
	helper.SetUpdated(&allowanceType.BaseDomain, user)

	return response.ToAllowanceTypeResponse(atService.AllowanceTypeRepository.Create(allowanceType))
}

func (atService *AllowanceTypeServiceImpl) Update(id int64, request *request.AllowanceTypeRequest, header dto.Header) response.AllowanceTypeResponse {
	atService.validateRequest(request)
	user := atService.UserService.FindByIdDomain(header.UserId)

	allowanceType := atService.FindByIdDomain(id)
	allowanceType.Name = request.Name
	helper.SetUpdated(&allowanceType.BaseDomain, user)

	return response.ToAllowanceTypeResponse(atService.AllowanceTypeRepository.Update(allowanceType))
}

func (atService *AllowanceTypeServiceImpl) FindById(id int64) response.AllowanceTypeResponse {
	allowanceType := atService.FindByIdDomain(id)
	return response.ToAllowanceTypeResponse(allowanceType)
}

func (atService *AllowanceTypeServiceImpl) FindByIdDomain(id int64) *domain.AllowanceType {
	allowanceType, err := atService.AllowanceTypeRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return allowanceType
}

func (atService *AllowanceTypeServiceImpl) FindAll(search *dto.Search) []response.AllowanceTypeResponse {
	allowanceTypes := atService.AllowanceTypeRepository.FindAll(search)
	return response.ToAllowanceTypeResponses(allowanceTypes)
}

func (atService *AllowanceTypeServiceImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse {
	allowanceTypes := atService.AllowanceTypeRepository.FindAllPagination(search, pagination)
	totalItem := atService.AllowanceTypeRepository.FindTotalItem()

	responses := response.ToAllowanceTypeResponses(allowanceTypes)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (atService *AllowanceTypeServiceImpl) Delete(id int64) response.AllowanceTypeResponse {
	allowanceType := atService.FindByIdDomain(id)
	atService.AllowanceTypeRepository.Delete(id)
	return response.ToAllowanceTypeResponse(allowanceType)
}

func (atService *AllowanceTypeServiceImpl) validateRequest(request *request.AllowanceTypeRequest) {
	err := atService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
}
