package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost  user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Нет подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&requestBody{}); err != nil {
		log.Fatalf("Мигрвция невозможна: %v", err)
	}
}

type requestBody struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}

//var tasks = []requestBody{}

func getTasks(w http.ResponseWriter) {
	var tasks []requestBody

	if err := db.Find(&tasks).Error; err != nil {
		http.Error(w, `не удалось получить таски`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)

}

func createTask(w http.ResponseWriter, r *http.Request) {
	var t requestBody

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, `Ошибка: неправильный JSON`, http.StatusBadRequest)
		return
	}
	if t.Task == "" {
		http.Error(w, `Ошибка: поле task пустое или не обнаружено`, http.StatusBadRequest)
		return
	}
	task := requestBody{
		ID:   uuid.NewString(),
		Task: t.Task,
	}

	//tasks = append(tasks, t)

	if err := db.Create(&task).Error; err != nil {
		http.Error(w, `не удалось создать таски`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/task/")
	var update requestBody

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, `Ошибка: неправильный JSON`, http.StatusBadRequest)
		return
	}

	if update.Task == "" {
		http.Error(w, `Ошибка: task пустой или не обнаружен`, http.StatusBadRequest)
		return
	}

	var task requestBody
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		http.Error(w, `Ошибка: выражение не найдено`, http.StatusBadRequest)
		return
	}
	task.Task = update.Task
	db.Save(&task)
	json.NewEncoder(w).Encode(task)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/task/")

	if err := db.Delete(&requestBody{}, "id = ?", id).Error; err != nil {
		http.Error(w, "не удалось удалить таск", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func taskHandles(w http.ResponseWriter, r *http.Request) {
	//path := r.URL.Path
	switch r.Method {
	case http.MethodGet:
		getTasks(w)
	case http.MethodPost:
		createTask(w, r)
	case http.MethodPatch:
		updateTask(w, r)
	case http.MethodDelete:
		deleteTask(w, r)

	default:
		http.Error(w, "метод недоступен", http.StatusMethodNotAllowed)
	}

}

func main() {
	initDB()
	http.HandleFunc("/task", taskHandles)
	http.HandleFunc("/task/", taskHandles)
	fmt.Println("Сервер запущен на localhost:8080")
	http.ListenAndServe(":8080", nil)
}
