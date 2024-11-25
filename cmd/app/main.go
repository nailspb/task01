package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"task01/internal/http/handlers"
	"task01/internal/services"
	"task01/internal/storage/postgres"
	"task01/internal/web/tasks"
	"task01/internal/web/users"
	"task01/pkg/prettylogger"
)

func main() {
	log := prettylogger.New("prod")
	log.Info("Start service")
	storage := postgres.New()
	taskHandlers := handlers.NewTasksHandler(services.NewTaskService(storage, log))
	userHandlers := handlers.NewUsersHandler(services.NewUserService(storage, log))

	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)
	e.GET("/api/*", echoSwagger.EchoWrapHandler(func(config *echoSwagger.Config) {
		config.URLs = []string{"/api/swagger.yaml"}
	}))

	// Сервим YAML-файл
	e.Static("/api/swagger.yaml", "./openapi/openapi.yaml")

	tasks.RegisterHandlers(e, tasks.NewStrictHandler(taskHandlers, nil))
	users.RegisterHandlers(e, users.NewStrictHandler(userHandlers, nil))

	if err := e.Start(":8080"); err != nil {
		panic(fmt.Errorf("failed to start with err: %v", err))
	}

}
