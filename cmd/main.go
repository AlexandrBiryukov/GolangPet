package main

import (
	"golang/internal/db"
	_ "golang/internal/db"
	Handlers "golang/internal/handlers"
	"golang/internal/services"
	"golang/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("нет подключения к бд: %v", err)
	}

	if err := database.AutoMigrate(&services.Tasks{}); err != nil {
		log.Fatalf("ошибка автомиграции: %v", err)
	}
	taskRepo := services.NewTaskRepositoty(database)
	taskService := services.NewTaskService(taskRepo)
	taskHandlels := Handlers.NewTaskHandler(taskService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.RequestLogger())

	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(taskHandlels, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
