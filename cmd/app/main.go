package main

import (
	"log"

	"restapi/internal/database"
	"restapi/internal/handlers"
	"restapi/internal/tasksService"
	"restapi/internal/web/tasks"
	"restapi/internal/web/users"
	"restapi/usersService"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	err := database.DB.AutoMigrate(&tasksService.Task{}, &usersService.User{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	taskRepo := tasksService.NewTaskRepository(database.DB)
	userRepo := usersService.NewUserRepository(database.DB)

	taskService := tasksService.NewService(taskRepo)
	userService := usersService.NewUserService(userRepo)

	handler := handlers.NewHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	users.RegisterHandlers(e, userHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
