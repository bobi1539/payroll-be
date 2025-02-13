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

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) controller.RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

// @Tags	Role
// @Accept	json
// @Produce	json
// @Param	request 		body	request.RoleRequest	true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.RoleResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/roles			[post]
func (roleController *RoleControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.RoleRequest](ctx)

	response := roleController.RoleService.Create(request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Role
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Param	request 		body	request.RoleRequest	true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.RoleResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/roles/id/{id}		[put]
func (roleController *RoleControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ROLE_ID)
	request := helper.BodyParser[request.RoleRequest](ctx)

	response := roleController.RoleService.Update(id, request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Role
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.RoleResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/roles/id/{id} 	[get]
func (roleController *RoleControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ROLE_ID)

	response := roleController.RoleService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Role
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Success	200	{object}	response.WebResponse{data=[]response.RoleResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/roles/all		[get]
func (roleController *RoleControllerImpl) FindAll(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	response := roleController.RoleService.FindAll(&search)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Role
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/roles			[get]
func (roleController *RoleControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := roleController.RoleService.FindAllPagination(&search, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Role
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.RoleResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/roles/id/{id} 	[delete]
func (roleController *RoleControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ROLE_ID)

	response := roleController.RoleService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
