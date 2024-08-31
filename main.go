package main

import (
	"log"

	"net/http"
	"verbatestapp/controllers"
	"verbatestapp/models"
)

func main() {
	err := models.OpenDb()
	if err != nil {
		log.Fatal("No connection to db:  ", err)
	}
	defer models.Db.Close()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetTasks(w, r)
		case http.MethodPost:
			controllers.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetTaskById(w, r)
		case http.MethodPut:
			controllers.UpdateTask(w, r)
		case http.MethodDelete:
			controllers.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is started!")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
