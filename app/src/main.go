package main

import (
	"app/controller"
	"app/db"
	"app/model"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Database
	db.Init()

	// Auto Migration
	if err := db.DB.AutoMigrate(&model.Sample{}); err != nil {
		slog.Error("failed to migrate database", "error", err)
	}

	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Initialize Controller
	sampleController := controller.SampleController{}

	// Routes
	router.GET("/", hello)
	router.GET("/hostname", sampleController.GetHostname)
	router.GET("/sample", sampleController.GetSample)
	router.POST("/sample", sampleController.PostSample)

	// Start server
	if err := router.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

// Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
