package usecases

import "template/pkg/todo"

// toDoUseCases to do usecases
type toDoUseCases struct {
	repo         todo.Repositories
	todoServicer todo.Services
}

// NewToDoUseCases create to do use case
func NewToDoUseCases(repo todo.Repositories, todoServicer todo.Services) todo.UseCases {
	return &toDoUseCases{
		repo:         repo,
		todoServicer: todoServicer,
	}
}

// CreateToDo create a to do
func (t *toDoUseCases) CreateToDo(todo todo.ToDo) (uint32, error) {
	if err := t.todoServicer.ValidateToDo(todo); err != nil {
		return 0, err
	}
	if err := t.repo.CreateToDo(&todo); err != nil {
		return 0, err
	}
	return todo.ID, nil
}
