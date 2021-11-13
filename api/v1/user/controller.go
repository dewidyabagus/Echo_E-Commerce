package user

import (
	echo "github.com/labstack/echo/v4"

	"RESTful/api/common"
	"RESTful/api/middleware"
	"RESTful/api/v1/user/request"
	"RESTful/api/v1/user/response"
	"RESTful/business/user"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) AddNewUser(ctx echo.Context) error {
	user := new(request.RequestUser)

	var err error
	err = ctx.Bind(user)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	id, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	err = c.service.AddNewUser(user.ToUserAddSpec(), &id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) FindUserByUserId(ctx echo.Context) error {
	id, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	user, err := c.service.FindUserByUserId(&id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetUserResponseDetail(user)))
}
