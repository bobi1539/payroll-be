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

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) controller.UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// @Tags	User
// @Accept	json
// @Produce	json
// @Param	request 		body	request.UserCreateRequest	true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.UserResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/users			[post]
// @Security 				BearerAuth
func (userController *UserControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.UserCreateRequest](ctx)

	response := userController.UserService.Create(request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	User
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Param	request 		body	request.UserUpdateRequest	true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.UserResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/users/id/{id}	[put]
// @Security 				BearerAuth
func (userController *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)
	request := helper.BodyParser[request.UserUpdateRequest](ctx)

	response := userController.UserService.Update(id, request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	User
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.UserResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/users/id/{id} 	[get]
// @Security 				BearerAuth
func (userController *UserControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := userController.UserService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	User
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Success	200	{object}	response.WebResponse{data=[]response.UserResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/users/all		[get]
// @Security 				BearerAuth
func (userController *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	response := userController.UserService.FindAll(&search)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	User
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/users			[get]
// @Security 				BearerAuth
func (userController *UserControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := userController.UserService.FindAllPagination(&search, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	User
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.UserResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/users/id/{id} 	[delete]
// @Security 				BearerAuth
func (userController *UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := userController.UserService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
