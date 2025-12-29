package controller

import (
	"app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SampleController struct {
	SampleService service.SampleService
}

func (c *SampleController) GetSample(ctx echo.Context) error {
	sample, err := c.SampleService.GetSample()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, sample)
}
