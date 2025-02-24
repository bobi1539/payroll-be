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

type BasicSalaryControllerImpl struct {
	BasicSalaryService service.BasicSalaryService
}

func NewBasicSalaryControllerImpl(basicSalaryService service.BasicSalaryService) controller.BasicSalaryController {
	return &BasicSalaryControllerImpl{BasicSalaryService: basicSalaryService}
}

// @Tags	BasicSalary
// @Accept	json
// @Produce	json
// @Param	request 		body	request.BasicSalaryRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.BasicSalaryResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/basic-salaries	[post]
// @Security 				BearerAuth
func (basicSalaryController *BasicSalaryControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.BasicSalaryRequest](ctx)
	header := dto.GetHeader(ctx)

	response := basicSalaryController.BasicSalaryService.Create(request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	BasicSalary
// @Accept	json
// @Produce	json
// @Param	id				path	int					true		"id"
// @Param	request 		body	request.BasicSalaryRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.BasicSalaryResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/basic-salaries/id/{id}	[put]
// @Security 				BearerAuth
func (basicSalaryController *BasicSalaryControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)
	request := helper.BodyParser[request.BasicSalaryRequest](ctx)
	header := dto.GetHeader(ctx)

	response := basicSalaryController.BasicSalaryService.Update(id, request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))

}

// @Tags	BasicSalary
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.BasicSalaryResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/basic-salaries/id/{id}	[get]
// @Security 				BearerAuth
func (basicSalaryController *BasicSalaryControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := basicSalaryController.BasicSalaryService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	BasicSalary
// @Accept	json
// @Produce	json
// @Param   positionId  	query	number	true	"Position Id"
// @Success	200	{object}	response.WebResponse{data=[]response.BasicSalaryResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/basic-salaries/all		[get]
// @Security 				BearerAuth
func (basicSalaryController *BasicSalaryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	bsSearch := search.BuildBasicSalarySearch(ctx.Query(constant.POSITION_ID))
	response := basicSalaryController.BasicSalaryService.FindAll(&bsSearch)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	BasicSalary
// @Accept	json
// @Produce	json
// @Param   positionId  	query	number	true	"Position Id"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/basic-salaries	[get]
// @Security 				BearerAuth
func (basicSalaryController *BasicSalaryControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	bsSearch := search.BuildBasicSalarySearch(ctx.Query(constant.POSITION_ID))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := basicSalaryController.BasicSalaryService.FindAllPagination(&bsSearch, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	BasicSalary
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.BasicSalaryResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/basic-salaries/id/{id} [delete]
// @Security 				BearerAuth
func (basicSalaryController *BasicSalaryControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := basicSalaryController.BasicSalaryService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
