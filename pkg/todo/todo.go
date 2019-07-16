package todo

import (
	"time"
)

// ToDo to do info
type ToDo struct {
	ID        uint32    `json:"id" gorm:"id" example:"123"`
	Title     string    `json:"title" gorm:"title" example:"todo title"`
	Content   string    `json:"content" gorm:"content" example:"todo content"`
	CreatedAt time.Time `gorm:"created_at" example:"2019-07-15T14:37:59Z"`
	UpdatedAt time.Time `gorm:"updated_at" example:"2019-07-15T14:37:59Z"`
}

// TableName todo table name
func (ToDo) TableName() string {
	return "todo"
}

// UseCases to do use case.
// Put specific application here
type UseCases interface {
	CreateToDo(todo ToDo) (uint32, error)
	ListToDos() ([]ToDo, error)
}

// Services to do services.
// Put common or core business logic here
type Services interface {
	ValidateToDo(todo ToDo) error
}

// Repositories to do repositories
type Repositories interface {
	CreateToDo(todo *ToDo) error
	ListToDos() ([]ToDo, error)
}
