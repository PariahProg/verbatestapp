/*Модель для взаимодействия с таблицей tasks*/

package models

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"verbatestapp/entities"
)

func CreateTask(task *entities.Task) error {
	task.CreatedAt = time.Now().Format(time.RFC3339)
	task.UpdatedAt = task.CreatedAt
	err := Db.QueryRow("INSERT INTO tasks (title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt).Scan(&task.Id)
	return err
}

func GetTasks() ([]entities.Task, error) {
	var tasks []entities.Task
	if rows, err := Db.Query("SELECT * FROM tasks"); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			t := entities.Task{}
			if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.DueDate, &t.CreatedAt, &t.UpdatedAt); err != nil {
				return nil, err
			}
			tasks = append(tasks, t)
		}
		return tasks, nil
	}
}

func GetTaskById(id int) (*entities.Task, error) {
	var task entities.Task
	err := Db.QueryRow("SELECT * FROM tasks WHERE id = $1", id).Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	} else {
		return &task, nil
	}
}

func UpdateTask(id int, task *entities.Task) error {
	task.Id = id
	task.UpdatedAt = time.Now().Format(time.RFC3339)
	err := Db.QueryRow("UPDATE tasks SET title = $1, description = $2, due_date = $3, updated_at = $4 WHERE id = $5 RETURNING created_at", task.Title, task.Description, task.DueDate, task.UpdatedAt, id).Scan(&task.CreatedAt)
	return err
}

func DeleteTask(id int) error {
	result, err := Db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error while deleting task: %w", err)
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while checking affecting rows: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("no task with provided id: %d", id)
	}
	return nil
}
