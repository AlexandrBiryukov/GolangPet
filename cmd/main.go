package main

import (
	"fmt"
	"golang/internal/db"
	Handlers "golang/internal/handlers"
	"golang/internal/services"
	"log"
	"net/http"
)

func main() {

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("нет подключения к бд: %v", err)
	}
	taskRepo := services.NewTaskRepositoty(database)
	taskService := services.NewTaskService(taskRepo)
	taskHandlels := Handlers.NewTaskHandler(taskService)

	http.HandleFunc("/task", taskHandlels.Handle)
	http.HandleFunc("/task/", taskHandlels.Handle)
	fmt.Println("Сервер запущен на localhost:8080")
	http.ListenAndServe(":8080", nil)
}
