package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"pokemon-tool-api/attribute/usecase"
)

// AttributeController struct
type AttributeController struct {
	attribute usecase.AttributeUsecase
}

// NewAttributeController handler
func NewAttributeController(
	e *echo.Echo,
	attributeU usecase.AttributeUsecase,
) {
	handler := &AttributeController{
		attribute: attributeU,
	}
	e.GET("/types", handler.List)
}

// List return typeList
func (a *AttributeController) List(ctx echo.Context) error {
	res, _ := a.attribute.List()
	return ctx.JSON(http.StatusOK, res)
}
