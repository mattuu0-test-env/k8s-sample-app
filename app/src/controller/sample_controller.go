package controller

import (
	"app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SampleController struct {
	SampleService service.SampleService
}

type CreateSampleRequest struct {
	Message string `json:"message"`
}

func (c *SampleController) GetSample(ctx echo.Context) error {
	sample, err := c.SampleService.GetSample()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, sample)
}

func (c *SampleController) PostSample(ctx echo.Context) error {
	req := new(CreateSampleRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if req.Message == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "message is required"})
	}

	sample, err := c.SampleService.CreateSample(req.Message)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, sample)
}