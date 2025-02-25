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

type BasicSalaryServiceImpl struct {
	BasicSalaryRepository repository.BasicSalaryRepository
	UserService           service.UserService
	PositionService       service.PositionService
	Validate              *validator.Validate
}

func NewBasicSalaryServiceImpl(
	basicSalaryRepository repository.BasicSalaryRepository,
	userService service.UserService,
	positionService service.PositionService,
	validate *validator.Validate,
) service.BasicSalaryService {
	return &BasicSalaryServiceImpl{
		BasicSalaryRepository: basicSalaryRepository,
		UserService:           userService,
		PositionService:       positionService,
		Validate:              validate,
	}
}

func (basicSalaryService *BasicSalaryServiceImpl) Create(request *request.BasicSalaryRequest, header dto.Header) response.BasicSalaryResponse {
	basicSalaryService.validateRequest(request)
	user := basicSalaryService.UserService.FindByIdDomain(header.UserId)

	basicSalary := &domain.BasicSalary{}
	basicSalaryService.setBasicSalary(basicSalary, request)
	helper.SetCreated(&basicSalary.BaseDomain, user)
	helper.SetUpdated(&basicSalary.BaseDomain, user)

	return response.ToBasicSalaryResponse(basicSalaryService.BasicSalaryRepository.Create(basicSalary))
}

func (basicSalaryService *BasicSalaryServiceImpl) Update(id int64, request *request.BasicSalaryRequest, header dto.Header) response.BasicSalaryResponse {
	basicSalaryService.validateRequest(request)
	user := basicSalaryService.UserService.FindByIdDomain(header.UserId)

	basicSalary := basicSalaryService.FindByIdDomain(id)
	basicSalaryService.setBasicSalary(basicSalary, request)
	helper.SetUpdated(&basicSalary.BaseDomain, user)

	return response.ToBasicSalaryResponse(basicSalaryService.BasicSalaryRepository.Update(basicSalary))
}

func (basicSalaryService *BasicSalaryServiceImpl) FindById(id int64) response.BasicSalaryResponse {
	basicSalary := basicSalaryService.FindByIdDomain(id)
	return response.ToBasicSalaryResponse(basicSalary)
}

func (basicSalaryService *BasicSalaryServiceImpl) FindByIdDomain(id int64) *domain.BasicSalary {
	basicSalary, err := basicSalaryService.BasicSalaryRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return basicSalary
}

func (basicSalaryService *BasicSalaryServiceImpl) FindAll(search *search.BasicSalarySearch) []response.BasicSalaryResponse {
	basicSalaries := basicSalaryService.BasicSalaryRepository.FindAll(search)
	return response.ToBasicSalaryResponses(basicSalaries)
}

func (basicSalaryService *BasicSalaryServiceImpl) FindAllPagination(search *search.BasicSalarySearch, pagination *dto.Pagination) response.PaginationResponse {
	basicSalaries := basicSalaryService.BasicSalaryRepository.FindAllPagination(search, pagination)
	totalItem := basicSalaryService.BasicSalaryRepository.FindTotalItem(search)

	responses := response.ToBasicSalaryResponses(basicSalaries)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (basicSalaryService *BasicSalaryServiceImpl) Delete(id int64) response.BasicSalaryResponse {
	basicSalary := basicSalaryService.FindByIdDomain(id)
	basicSalaryService.BasicSalaryRepository.Delete(id)
	return response.ToBasicSalaryResponse(basicSalary)
}

func (basicSalaryService *BasicSalaryServiceImpl) validateRequest(request *request.BasicSalaryRequest) {
	err := basicSalaryService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
}

func (basicSalaryService *BasicSalaryServiceImpl) setBasicSalary(basicSalary *domain.BasicSalary, request *request.BasicSalaryRequest) {
	basicSalary.SalaryAmount = request.SalaryAmount
	basicSalary.TotalYear = request.TotalYear
	basicSalary.Position = basicSalaryService.PositionService.FindByIdDomain(request.PositionId)
}
