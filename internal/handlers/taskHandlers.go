package Handlers

import (
	"encoding/json"
	"golang/internal/services"
	"net/http"
	"strings"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(s services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		http.Error(w, `не удалось получить таски`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)

}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t services.RequestBody

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, `Ошибка: неправильный JSON`, http.StatusBadRequest)
		return
	}
	if t.Task == "" {
		http.Error(w, `Ошибка: поле task пустое или не обнаружено`, http.StatusBadRequest)
		return
	}
	task, err := h.service.CreateTask(t.Task)
	if err != nil {
		http.Error(w, `не удалось создать таски`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/task/")
	if id == "" {
		http.Error(w, "id пустой", http.StatusBadRequest)
		return
	}

	var update services.RequestBody

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, `Ошибка: неправильный JSON`, http.StatusBadRequest)
		return
	}

	if update.Task == "" {
		http.Error(w, `Ошибка: task пустой или не обнаружен`, http.StatusBadRequest)
		return
	}

	updatedTask, err := h.service.UpdateTask(id, update.Task)
	if err != nil {
		http.Error(w, `не удалось обновить таски`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)

}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/task/")

	if err := h.service.DeleteTask(id); err != nil {
		http.Error(w, "не удалось удалить таск", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func (h *TaskHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTasks(w)
	case http.MethodPost:
		h.CreateTask(w, r)
	case http.MethodPatch:
		h.UpdateTask(w, r)
	case http.MethodDelete:
		h.DeleteTask(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
