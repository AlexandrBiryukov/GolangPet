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
		json.NewDecoder(r.Body).Decode(&body)
		task = body.Task
		return
	}

	if r.Method == http.MethodGet {
		w.Write(([]byte)("hello, " + task))
		return
	}

}
