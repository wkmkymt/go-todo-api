package datastore

import "github.com/wkmkymt/go-todo-api/domain/model"

type todoRepository struct {
}

// TodoRepository is Todo Repository
type TodoRepository interface {
	FindAll() (model.Todos, error)
}

// NewTodoRepository is Creating New Todo Repository
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (ur *todoRepository) FindAll() (model.Todos, error) {
	todos := model.Todos{}

	return todos, nil
}
