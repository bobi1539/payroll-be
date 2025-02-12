package controllerimpl

import (
	"payroll/constant"
	"payroll/controller"
	"payroll/helper"
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

func (roleController *RoleControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.RoleRequest](ctx)

	response := roleController.RoleService.Create(request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

func (roleController *RoleControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ROLE_ID)
	request := helper.BodyParser[request.RoleRequest](ctx)

	response := roleController.RoleService.Update(id, request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

func (roleController *RoleControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ROLE_ID)

	response := roleController.RoleService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
