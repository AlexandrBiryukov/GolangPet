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
		var body requestBody

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Неправильный json", http.StatusBadRequest)
			return
		}
		task = body.Task

		w.Write([]byte("task обновлен"))
		return
	}

	if r.Method == http.MethodGet {
		w.Write(([]byte)("hello, " + task))
		return
	}

}
