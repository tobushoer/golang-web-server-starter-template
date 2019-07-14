package di

import (
	"template/internal/pkg/database/mysql"
	"template/pkg/todo"
	"template/pkg/todo/repositories"
	"template/pkg/todo/services"
	"template/pkg/todo/usecases"
)

// NewToDoUseCases create todo use cases instance
func NewToDoUseCases() todo.UseCases {
	db := mysql.DB()
	toDoRepo := repositories.NewRepositories(db)
	toDoService := services.NewToDoServicer()
	return usecases.NewToDoUseCases(toDoRepo, toDoService)
}
