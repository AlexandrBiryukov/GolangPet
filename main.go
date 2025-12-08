package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type requestBody struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var tasks = []requestBody{}

func handleTask(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		body, err := validJSON(r, false, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task := requestBody{
			ID:   uuid.NewString(),
			Task: body.Task,
		}
		tasks = append(tasks, task)
		w.Write(([]byte)("Запрос выполнен"))
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
		return
	}

	if r.Method == http.MethodPatch {

		body, err := validJSON(r, true, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for id, v := range tasks {
			if v.ID == body.ID {
				tasks[id].Task = body.Task
				break
			}

		}
		w.Write([]byte("Изменено"))
		return

	}

	if r.Method == http.MethodDelete {
		body, err := validJSON(r, true, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for id, v := range tasks {
			if v.ID == body.ID {
				tasks = append(tasks[:id], tasks[id+1:]...)
				break
			}

		}
		w.Write([]byte("Удалено"))
		return
	}

}

func validJSON(r *http.Request, ID bool, Task bool) (requestBody, error) {
	var raw map[string]interface{}
	var body requestBody

	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		return body, fmt.Errorf("Неправильный json")
	}

	if ID {
		if _, ok := raw["id"]; !ok {
			return body, fmt.Errorf("Нет поля id")
		}
	}

	if Task {
		if _, ok := raw["task"]; !ok {
			return body, fmt.Errorf("Нет поля task")
		}
	}

	rawBytes, _ := json.Marshal(raw)
	json.Unmarshal(rawBytes, &body)

	return body, nil
}

func main() {

	http.HandleFunc("/task", handleTask)
	fmt.Println("Сервер запущен на localhost:8080")
	http.ListenAndServe(":8080", nil)
}
