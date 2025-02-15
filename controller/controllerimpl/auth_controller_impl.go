package controllerimpl

import (
	"payroll/controller"
	"payroll/helper"
	"payroll/model/request"
	"payroll/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) controller.AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	request 		body	request.LoginRequest	true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.LoginResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/auths/login	[post]
func (authController *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.LoginRequest](ctx)

	response := authController.AuthService.Login(request)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
