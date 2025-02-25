package serviceimpl

import (
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/model/search"
	"payroll/repository"
	"payroll/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AllowanceServiceImpl struct {
	AllowanceRepository  repository.AllowanceRepository
	UserService          service.UserService
	PositionService      service.PositionService
	AllowanceTypeService service.AllowanceTypeService
	Validate             *validator.Validate
}

func NewAllowanceServiceImpl(
	allowanceRepository repository.AllowanceRepository,
	userService service.UserService,
	positionService service.PositionService,
	allowanceTypeService service.AllowanceTypeService,
	validate *validator.Validate,
) service.AllowanceService {
	return &AllowanceServiceImpl{
		AllowanceRepository:  allowanceRepository,
		UserService:          userService,
		PositionService:      positionService,
		AllowanceTypeService: allowanceTypeService,
		Validate:             validate,
	}
}

func (allowanceService *AllowanceServiceImpl) Create(request *request.AllowanceRequest, header dto.Header) response.AllowanceResponse {
	allowanceService.validateRequest(request)
	user := allowanceService.UserService.FindByIdDomain(header.UserId)

	allowance := &domain.Allowance{}
	allowanceService.setAllowance(allowance, request)
	helper.SetCreated(&allowance.BaseDomain, user)
	helper.SetUpdated(&allowance.BaseDomain, user)

	return response.ToAllowanceResponse(allowanceService.AllowanceRepository.Create(allowance))
}

func (allowanceService *AllowanceServiceImpl) Update(id int64, request *request.AllowanceRequest, header dto.Header) response.AllowanceResponse {
	allowanceService.validateRequest(request)
	user := allowanceService.UserService.FindByIdDomain(header.UserId)

	allowance := allowanceService.FindByIdDomain(id)
	allowanceService.setAllowance(allowance, request)
	helper.SetUpdated(&allowance.BaseDomain, user)

	return response.ToAllowanceResponse(allowanceService.AllowanceRepository.Update(allowance))
}

func (allowanceService *AllowanceServiceImpl) FindById(id int64) response.AllowanceResponse {
	allowance := allowanceService.FindByIdDomain(id)
	return response.ToAllowanceResponse(allowance)
}

func (allowanceService *AllowanceServiceImpl) FindByIdDomain(id int64) *domain.Allowance {
	allowance, err := allowanceService.AllowanceRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return allowance
}

func (allowanceService *AllowanceServiceImpl) FindAll(search *search.AllowanceSearch) []response.AllowanceResponse {
	allowances := allowanceService.AllowanceRepository.FindAll(search)
	return response.ToAllowanceResponses(allowances)
}

func (allowanceService *AllowanceServiceImpl) FindAllPagination(search *search.AllowanceSearch, pagination *dto.Pagination) response.PaginationResponse {
	allowances := allowanceService.AllowanceRepository.FindAllPagination(search, pagination)
	totalItem := allowanceService.AllowanceRepository.FindTotalItem(search)

	responses := response.ToAllowanceResponses(allowances)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (allowanceService *AllowanceServiceImpl) Delete(id int64) response.AllowanceResponse {
	allowance := allowanceService.FindByIdDomain(id)
	allowanceService.AllowanceRepository.Delete(id)
	return response.ToAllowanceResponse(allowance)
}

func (allowanceService *AllowanceServiceImpl) validateRequest(request *request.AllowanceRequest) {
	err := allowanceService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
}

func (allowanceService *AllowanceServiceImpl) setAllowance(allowance *domain.Allowance, request *request.AllowanceRequest) {
	allowance.Position = allowanceService.PositionService.FindByIdDomain(request.PositionId)
	allowance.AllowanceType = allowanceService.AllowanceTypeService.FindByIdDomain(request.AllowanceTypeId)
	allowance.AllowanceAmount = request.AllowanceAmount
}
