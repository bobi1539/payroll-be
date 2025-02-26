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

type EmployeeControllerImpl struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeControllerImpl(employeeService service.EmployeeService) controller.EmployeeController {
	return &EmployeeControllerImpl{EmployeeService: employeeService}
}

// @Tags	Employee
// @Accept	json
// @Produce	json
// @Param	request 		body	request.EmployeeRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.EmployeeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/employees		[post]
// @Security 				BearerAuth
func (employeeController *EmployeeControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.EmployeeRequest](ctx)
	header := dto.GetHeader(ctx)

	response := employeeController.EmployeeService.Create(request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Employee
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Param	request 		body	request.EmployeeRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.EmployeeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/employees/id/{id}		[put]
// @Security 				BearerAuth
func (employeeController *EmployeeControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)
	request := helper.BodyParser[request.EmployeeRequest](ctx)
	header := dto.GetHeader(ctx)

	response := employeeController.EmployeeService.Update(id, request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Employee
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.EmployeeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/employees/id/{id} 		[get]
// @Security 				BearerAuth
func (employeeController *EmployeeControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := employeeController.EmployeeService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Employee
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Success	200	{object}	response.WebResponse{data=[]response.EmployeeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/employees/all	[get]
// @Security 				BearerAuth
func (employeeController *EmployeeControllerImpl) FindAll(ctx *fiber.Ctx) error {
	search := search.BuildEmployeeSearch(ctx.Query(constant.SEARCH))

	response := employeeController.EmployeeService.FindAll(&search)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Employee
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/employees		[get]
// @Security 				BearerAuth
func (employeeController *EmployeeControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	search := search.BuildEmployeeSearch(ctx.Query(constant.SEARCH))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := employeeController.EmployeeService.FindAllPagination(&search, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Employee
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.EmployeeResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/employees/id/{id} 		[delete]
// @Security 				BearerAuth
func (employeeController *EmployeeControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := employeeController.EmployeeService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
