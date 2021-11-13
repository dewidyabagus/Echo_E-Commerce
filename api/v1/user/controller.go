package user

import (
	"github.com/google/uuid"
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
	var err error

	adminId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	user := new(request.RequestUser)
	if err = ctx.Bind(user); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err = c.service.AddNewUser(user.ToUserAddSpec(), &adminId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) FindAllUser(ctx echo.Context) error {
	adminId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	users, err := c.service.FindAllUser(&adminId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetAllUserResponseDetail(users)))
}

func (c *Controller) FindUserByUserId(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	adminId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	user, err := c.service.FindUserByUserId(&id, &adminId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetUserResponseDetail(user)))
}

func (c *Controller) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	adminId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	user := new(request.RequestUser)
	if err = ctx.Bind(user); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.UpdateUser(&id, user.ToUserUpateSpec(), &adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	adminId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.DeleteUser(&id, &adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
