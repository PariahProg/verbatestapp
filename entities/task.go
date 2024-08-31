/*Пакет с кастомными типами данных и структурами. В данном проекте необходима только структура Task, которая хранит в себе данные конкретной задачи. */

package entities

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
