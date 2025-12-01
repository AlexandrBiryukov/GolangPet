package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

func main() {

	http.HandleFunc("/task", handleTask)
	fmt.Println("Сервер запущен на localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type requestBody struct {
	Task string `json:"task"`
}

func handleTask(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var raw map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&raw)
		if err != nil {
			http.Error(w, "Неправильный json", http.StatusBadRequest)
			return
		}

		_, ok := raw["task"]
		if !ok {
			http.Error(w, "Нет поля task", http.StatusBadRequest)
			return
		}

		var body requestBody
		rawBytes, _ := json.Marshal(raw)
		json.Unmarshal(rawBytes, &body)

		w.Write(([]byte)("Запрос выполнен"))
		task = body.Task
		return
	}

	if r.Method == http.MethodGet {
		w.Write(([]byte)("Hello, " + task))
		return
	}

}
