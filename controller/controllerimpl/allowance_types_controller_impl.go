package controllerimpl

import (
	"payroll/constant"
	"payroll/controller"
	"payroll/helper"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/service"

	"github.com/gofiber/fiber/v2"
)

type AllowanceTypeControllerImpl struct {
	AllowanceTypeService service.AllowanceTypeService
}

func NewAllowanceTypeControllerImpl(allowanceTypeService service.AllowanceTypeService) controller.AllowanceTypeController {
	return &AllowanceTypeControllerImpl{AllowanceTypeService: allowanceTypeService}
}

// @Tags	AllowanceType
// @Accept	json
// @Produce	json
// @Param	request 		body	request.AllowanceTypeRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceTypeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowance-types		[post]
// @Security 				BearerAuth
func (atController *AllowanceTypeControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.AllowanceTypeRequest](ctx)
	header := dto.GetHeader(ctx)

	response := atController.AllowanceTypeService.Create(request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	AllowanceType
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Param	request 		body	request.AllowanceTypeRequest	true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceTypeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowance-types/id/{id}					[put]
// @Security 				BearerAuth
func (atController *AllowanceTypeControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)
	request := helper.BodyParser[request.AllowanceTypeRequest](ctx)
	header := dto.GetHeader(ctx)

	response := atController.AllowanceTypeService.Update(id, request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	AllowanceType
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceTypeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowance-types/id/{id} 					[get]
// @Security 				BearerAuth
func (atController *AllowanceTypeControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := atController.AllowanceTypeService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	AllowanceType
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Success	200	{object}	response.WebResponse{data=[]response.AllowanceTypeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowance-types/all	[get]
// @Security 				BearerAuth
func (atController *AllowanceTypeControllerImpl) FindAll(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	response := atController.AllowanceTypeService.FindAll(&search)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	AllowanceType
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowance-types				[get]
// @Security 				BearerAuth
func (atController *AllowanceTypeControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := atController.AllowanceTypeService.FindAllPagination(&search, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	AllowanceType
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceTypeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowance-types/id/{id} 					[delete]
// @Security 				BearerAuth
func (atController *AllowanceTypeControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := atController.AllowanceTypeService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
