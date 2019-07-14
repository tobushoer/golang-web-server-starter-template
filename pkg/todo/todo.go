package todo

import (
	"time"
)

// ToDo to do info
type ToDo struct {
	ID        uint32    `json:"id" gorm:"id"`
	Title     string    `json:"title" gorm:"title"`
	Content   string    `json:"content" gorm:"content"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

// TableName todo table name
func (ToDo) TableName() string {
	return "todo"
}

// UseCases to do use case.
// Put specific application here
type UseCases interface {
	CreateToDo(todo ToDo) (uint32, error)
}

// Services to do services.
// Put common or core business logic here
type Services interface {
	ValidateToDo(todo ToDo) error
}

// Repositories to do repositories
type Repositories interface {
	CreateToDo(todo *ToDo) error
}
