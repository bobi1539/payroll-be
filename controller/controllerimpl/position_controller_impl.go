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

type PositionControllerImpl struct {
	PositionService service.PositionService
}

func NewPositionControllerImpl(positionService service.PositionService) controller.PositionController {
	return &PositionControllerImpl{
		PositionService: positionService,
	}
}

// @Tags	Position
// @Accept	json
// @Produce	json
// @Param	request 		body	request.PositionRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.PositionResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/positions		[post]
// @Security 				BearerAuth
func (positionController *PositionControllerImpl) Create(ctx *fiber.Ctx) error {
	request := helper.BodyParser[request.PositionRequest](ctx)
	header := dto.GetHeader(ctx)

	response := positionController.PositionService.Create(request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Position
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Param	request 		body	request.PositionRequest		true	"Request body"
// @Success	200	{object}	response.WebResponse{data=response.PositionResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/positions/id/{id}		[put]
// @Security 				BearerAuth
func (positionController *PositionControllerImpl) Update(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)
	request := helper.BodyParser[request.PositionRequest](ctx)
	header := dto.GetHeader(ctx)

	response := positionController.PositionService.Update(id, request, header)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Position
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.PositionResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/positions/id/{id} 		[get]
// @Security 				BearerAuth
func (positionController *PositionControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := positionController.PositionService.FindById(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Position
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Success	200	{object}	response.WebResponse{data=[]response.PositionResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/positions/all	[get]
// @Security 				BearerAuth
func (positionController *PositionControllerImpl) FindAll(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	response := positionController.PositionService.FindAll(&search)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Position
// @Accept	json
// @Produce	json
// @Param   search  		query	string	false	"Search"
// @Param   pageNumber  	query	string	false	"Page Number"	default(1)
// @Param   pageSize  		query	string	false	"Page Size"		default(10)
// @Success	200	{object}	response.WebResponse{data=[]response.PaginationResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/positions		[get]
// @Security 				BearerAuth
func (positionController *PositionControllerImpl) FindAllPagination(ctx *fiber.Ctx) error {
	search := dto.BuildSearch(ctx.Query(constant.SEARCH))

	pageNumber := ctx.Query(constant.PAGE_NUMBER)
	pageSize := ctx.Query(constant.PAGE_SIZE)
	pagination := dto.BuildPagination(pageNumber, pageSize)

	response := positionController.PositionService.FindAllPagination(&search, &pagination)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}

// @Tags	Position
// @Accept	json
// @Produce	json
// @Param	id				path	int					true	"id"
// @Success	200	{object}	response.WebResponse{data=response.PositionResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/positions/id/{id} 		[delete]
// @Security 				BearerAuth
func (positionController *PositionControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := helper.GetParamId(ctx, constant.ID)

	response := positionController.PositionService.Delete(id)
	return ctx.JSON(helper.BuildSuccessResponse(response))
}
