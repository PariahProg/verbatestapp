/* Контроллер, обрабатывающий все эндпоинты данного API.
Сервер возвращает все HTTP-коды, указанные в задании, согласно параметрам запроса.
От себя я добавил код 415, если у запроса неверный Content-Type, и 400, если неверный формат id (например отрицательное число).*/

package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"verbatestapp/entities"
	"verbatestapp/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) { // Обработчик для POST /tasks
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "Неподдерживаемый content type!", http.StatusUnsupportedMediaType)
		return
	}

	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Неправильный формат данных!", http.StatusBadRequest)
		return
	}

	if task.Title == "" || task.Description == "" || task.DueDate == "" {
		http.Error(w, "Неправильный формат данных!", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse(time.RFC3339, task.DueDate); err != nil {
		http.Error(w, "Неправильный формат данных!", http.StatusBadRequest)
		return
	}

	if err := models.CreateTask(&task); err != nil {
		log.Println(err)
		http.Error(w, "Проблема на сервере!", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&task)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) { // Обработчик для GET /tasks
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "Неподдерживаемый content type!", http.StatusUnsupportedMediaType)
		return
	}

	if tasks, err := models.GetTasks(); err != nil {
		log.Println(err)
		http.Error(w, "Внутренняя ошибка сервера!", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) { // Обработчик для GET /tasks/{id}
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "Неподдерживаемый content type!", http.StatusUnsupportedMediaType)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "Неверный формат id!", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		http.Error(w, "Неверный формат id!", http.StatusBadRequest)
		return
	}

	if task, err := models.GetTaskById(id); err != nil && errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Задача не найдена!", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Проблема на сервере!", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&task)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) { // Обработчик для PUT /tasks/{id}
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "Неподдерживаемый content type!", http.StatusUnsupportedMediaType)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "Неверный формат id!", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		http.Error(w, "Неверный формат id!", http.StatusBadRequest)
		return
	}

	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Неправильный формат данных!", http.StatusBadRequest)
		return
	}

	if task.Title == "" || task.Description == "" || task.DueDate == "" {
		http.Error(w, "Неправильный формат данных!", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse(time.RFC3339, task.DueDate); err != nil {
		http.Error(w, "Неправильный формат данных!", http.StatusBadRequest)
		return
	}

	if err := models.UpdateTask(id, &task); err != nil && errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Задача не найдена!", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Проблема на сервере!", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&task)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) { // Обработчик для DELETE /tasks/{id}
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "Неподдерживаемый content type!", http.StatusUnsupportedMediaType)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "Неверный формат id!", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		http.Error(w, "Неверный формат id!", http.StatusBadRequest)
		return
	}

	if err := models.DeleteTask(id); err != nil && err.Error() == fmt.Sprintf("no task with provided id: %d", id) {
		http.Error(w, "Задача не найдена!", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		http.Error(w, "Проблема на сервере!", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
