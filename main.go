package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	u "github.com/Yandex-Practicum/go-rest-api-homework/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	res, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.SetUpHeaders(http.StatusOK, w)
	w.Write(res)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	res, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	u.SetUpHeaders(http.StatusOK, w)
	w.Write(res)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tasks[task.ID] = task

	u.SetUpHeaders(http.StatusCreated, w)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	delete(tasks, id)

	u.SetUpHeaders(http.StatusOK, w)
}

func main() {

	r := chi.NewRouter()

	r.Get("/tasks", GetTasks)
	r.Get("/tasks/{id}", GetTask)
	r.Post("/tasks", PostTask)
	r.Delete("/tasks/{id}", DeleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
