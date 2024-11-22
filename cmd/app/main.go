package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task01/internal/http/handlers"
	"task01/internal/services"
	"task01/internal/storage/postgres"
	"task01/internal/web/tasks"
	"task01/pkg/prettylogger"
)

func main() {
	log := prettylogger.New("prod")
	log.Info("Start service")
	taskHandlers := handlers.NewTasksHandler(services.NewTaskService(postgres.New(), log), log)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)
	if err := e.Start(":8080"); err != nil {
		panic(fmt.Errorf("failed to start with err: %v", err))
	}

}
