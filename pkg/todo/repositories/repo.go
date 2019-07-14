package repositories

import (
	"template/internal/pkg/log"
	"template/pkg/todo"

	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

// NewRepositories create todo repository instance
func NewRepositories(db *gorm.DB) todo.Repositories {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateToDo(toDo *todo.ToDo) error {
	if err := r.db.Model(&todo.ToDo{}).Create(toDo).Error; err != nil {
		log.Errorf("create new to do failed: %s", err.Error())
		return err
	}
	return nil
}
