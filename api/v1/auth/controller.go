package auth

import (
	"RESTful/api/common"
	"RESTful/api/v1/auth/request"
	"RESTful/api/v1/auth/response"
	"RESTful/business/auth"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service auth.Service
}

func NewController(service auth.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) LoginUserWithEmailPassword(ctx echo.Context) error {
	var user = new(request.LoginRequest)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	token, err := c.service.UserLoginWithEmailPassword(user.ToLoginSpec())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	var loginResponse = response.SuccessLogin{
		Token: *token,
	}

	return ctx.JSON(common.SuccessResponseWithData(loginResponse))
}

func (c *Controller) LoginAdminWithEmailPassword(ctx echo.Context) error {
	var user = new(request.LoginRequest)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	token, err := c.service.AdminLoginWithEmailPassword(user.ToLoginSpec())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	var loginResponse = response.SuccessLogin{
		Token: *token,
	}

	return ctx.JSON(common.SuccessResponseWithData(loginResponse))
}
