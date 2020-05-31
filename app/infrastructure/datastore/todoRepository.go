package datastore

import (
	"github.com/jinzhu/gorm"

	"github.com/wkmkymt/go-todo-api/domain/model"
)

type todoRepository struct {
	db *gorm.DB
}

// TodoRepository is Todo Repository
type TodoRepository interface {
	Store(todo *model.Todo) bool
	FindAll() (model.Todos, error)
}

// NewTodoRepository is Creating New Todo Repository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (ur *todoRepository) Store(todo *model.Todo) bool {
	ur.db.NewRecord(todo)
	ur.db.Create(&todo)

	return ur.db.NewRecord(todo)
}

func (ur *todoRepository) FindAll() (model.Todos, error) {
	todos := model.Todos{}
	err := ur.db.Find(&todos).Error

	if err != nil {
		return nil, err
	}

	return todos, nil
}
