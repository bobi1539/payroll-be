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

type PositionServiceImpl struct {
	PositionRepository repository.PositionRepository
	UserService        service.UserService
	Validate           *validator.Validate
}

func NewPositionServiceImpl(
	positionRepository repository.PositionRepository,
	userService service.UserService,
	validate *validator.Validate,
) service.PositionService {
	return &PositionServiceImpl{
		PositionRepository: positionRepository,
		UserService:        userService,
		Validate:           validate,
	}
}

func (positionService *PositionServiceImpl) Create(request *request.PositionRequest, header dto.Header) response.PositionResponse {
	positionService.validateRequest(request)
	user := positionService.UserService.FindByIdDomain(header.UserId)

	position := &domain.Position{}
	position.Name = request.Name
	helper.SetCreated(&position.BaseDomain, user)
	helper.SetUpdated(&position.BaseDomain, user)

	return response.ToPositionResponse(positionService.PositionRepository.Create(position))
}

func (positionService *PositionServiceImpl) Update(id int64, request *request.PositionRequest, header dto.Header) response.PositionResponse {
	positionService.validateRequest(request)
	user := positionService.UserService.FindByIdDomain(header.UserId)

	position := positionService.FindByIdDomain(id)
	position.Name = request.Name
	helper.SetUpdated(&position.BaseDomain, user)

	return response.ToPositionResponse(positionService.PositionRepository.Update(position))
}

func (positionService *PositionServiceImpl) FindById(id int64) response.PositionResponse {
	position := positionService.FindByIdDomain(id)
	return response.ToPositionResponse(position)
}

func (positionService *PositionServiceImpl) FindByIdDomain(id int64) *domain.Position {
	position, err := positionService.PositionRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return position
}

func (positionService *PositionServiceImpl) FindAll(search *dto.Search) []response.PositionResponse {
	positions := positionService.PositionRepository.FindAll(search)
	return response.ToPositionResponses(positions)
}

func (positionService *PositionServiceImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse {
	positions := positionService.PositionRepository.FindAllPagination(search, pagination)
	totalItem := positionService.PositionRepository.FindTotalItem()

	responses := response.ToPositionResponses(positions)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (positionService *PositionServiceImpl) Delete(id int64) response.PositionResponse {
	position := positionService.FindByIdDomain(id)
	positionService.PositionRepository.Delete(id)
	return response.ToPositionResponse(position)
}

func (positionService *PositionServiceImpl) validateRequest(request *request.PositionRequest) {
	err := positionService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
}
