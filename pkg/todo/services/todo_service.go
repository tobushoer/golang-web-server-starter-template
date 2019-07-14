package services

import (
	"template/internal/pkg/errordef"
	"template/pkg/todo"
)

// ToDoServices to do service
type ToDoServices struct{}

// NewToDoServicer create to do servicer
func NewToDoServicer() todo.Services {
	return &ToDoServices{}
}

// ValidateToDo check if a to do valid
func (t *ToDoServices) ValidateToDo(task todo.ToDo) error {
	if task.Title == "" {
		return errordef.ErrMsgInvalidRequest
	}
	return nil
}
