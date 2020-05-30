package datastore

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/wkmkymt/go-todo-api/domain/model"
)

type todoRepository struct {
	db *gorm.DB
}

// TodoRepository is Todo Repository
type TodoRepository interface {
	FindAll() (model.Todos, error)
}

// NewTodoRepository is Creating New Todo Repository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (ur *todoRepository) FindAll() (model.Todos, error) {
	todos := model.Todos{
		{
			ID:        "1",
			Title:     "Todo 1",
			Status:    true,
			LimitDate: time.Now(),
		},
		{
			ID:        "2",
			Title:     "Todo 2",
			Status:    false,
			LimitDate: time.Now(),
		},
	}

	return todos, nil
}
