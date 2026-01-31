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
	"golang/internal/userService"
	"golang/internal/web/users"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("нет подключения к бд: %v", err)
	}

	//if err := database.AutoMigrate(&services.Tasks{}, &userService.User{}); err != nil {
	//	log.Fatalf("ошибка автомиграции: %v", err)
	//}
	taskRepo := services.NewTaskRepositoty(database)
	taskService := services.NewTaskService(taskRepo)
	taskHandlers := Handlers.NewTaskHandler(taskService)

	usersRepo := userService.NewRepository(database)
	usersService := userService.NewService(usersRepo)
	usersHandlers := Handlers.NewUserHandler(usersService)

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	tasksStrict := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, tasksStrict)

	usersStrict := users.NewStrictHandler(usersHandlers, nil)
	users.RegisterHandlers(e, usersStrict)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Не удалось запустить из-за ошибки: %v", err)
	}
}
