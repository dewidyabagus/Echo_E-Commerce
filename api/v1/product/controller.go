package product

import (
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"RESTful/api/common"
	"RESTful/api/middleware"
	"RESTful/api/v1/product/request"
	"RESTful/api/v1/product/response"
	"RESTful/business/product"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) FindAllProduct(ctx echo.Context) error {
	userId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	products, err := c.service.FindAllProduct(&userId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(
		response.GetAllProductDetail(products),
	))
}

func (c *Controller) AddNewProduct(ctx echo.Context) error {
	userId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	var product = new(request.RequestProduct)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.AddNewProduct(product.ToProductSpec(), &userId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) FindProductById(ctx echo.Context) error {
	userId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	product, err := c.service.FindProductById(&id, &userId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetProductDetail(product)))
}

func (c *Controller) UpdateProduct(ctx echo.Context) error {
	userId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	var product = new(request.RequestProduct)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.UpdateProduct(&id, product.ToProductSpec(), &userId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteProduct(ctx echo.Context) error {
	userId, err := middleware.ExtractJWTUserId(ctx)
	if err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.DeleteProduct(&id, &userId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
