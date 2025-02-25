package controllerimpl

import (
	"payroll/constant"
	"payroll/controller"
	"payroll/helper"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/search"
	"payroll/service"

	"github.com/gofiber/fiber/v2"
)

type AllowanceControllerImpl struct {
	AllowanceService service.AllowanceService
}

func NewAllowanceControllerImpl(allowanceService service.AllowanceService) controller.AllowanceController {
	return &AllowanceControllerImpl{AllowanceService: allowanceService}
}

// @Tags	Allowance
// @Accept	json
// @Produce	json
// @Param	request 		body	request.AllowanceRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowances		[post]
// @Security 				BearerAuth
func (allowanceController *AllowanceControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.AllowanceRequest](ctx)
	header := dto.GetHeader(ctx)

	response := allowanceController.AllowanceService.Create(request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Allowance
// @Accept	json
// @Produce	json
// @Param	id				path	int					true		"id"
// @Param	request 		body	request.AllowanceRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowances/id/{id}		[put]
// @Security 				BearerAuth
func (allowanceController *AllowanceControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)
	request := helper.BodyParser[request.AllowanceRequest](ctx)
	header := dto.GetHeader(ctx)

	response := allowanceController.AllowanceService.Update(id, request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))

}

// @Tags	Allowance
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowances/id/{id}		[get]
// @Security 				BearerAuth
func (allowanceController *AllowanceControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := allowanceController.AllowanceService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Allowance
// @Accept	json
// @Produce	json
// @Param   positionId  	query	number	true	"Position Id"
// @Success	200	{object}	response.WebResponse{data=[]response.AllowanceResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowances/all	[get]
// @Security 				BearerAuth
func (allowanceController *AllowanceControllerImpl) FindAll(ctx *fiber.Ctx) error {
	aSearch := search.BuildAllowanceSearch(ctx.Query(constant.POSITION_ID))
	response := allowanceController.AllowanceService.FindAll(&aSearch)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Allowance
// @Accept	json
// @Produce	json
// @Param   positionId  	query	number	true	"Position Id"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowances		[get]
// @Security 				BearerAuth
func (allowanceController *AllowanceControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	aSearch := search.BuildAllowanceSearch(ctx.Query(constant.POSITION_ID))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := allowanceController.AllowanceService.FindAllPagination(&aSearch, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Allowance
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.AllowanceResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/allowances/id/{id} 	[delete]
// @Security 				BearerAuth
func (allowanceController *AllowanceControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := allowanceController.AllowanceService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
