package admin

import (
	echo "github.com/labstack/echo/v4"

	"RESTful/api/common"
	"RESTful/api/v1/admin/request"
	"RESTful/business/admin"
)

type Controller struct {
	service admin.Service
}

func NewController(service admin.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) AddNewAdmin(ctx echo.Context) error {
	adminData := new(request.RequestAdmin)

	if err := ctx.Bind(adminData); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.AddNewAdmin(adminData.ToAdminAddSpec())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
