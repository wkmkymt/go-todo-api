package repository

import "github.com/wkmkymt/go-todo-api/domain/model"

// TodoRepository is Todo Repository
type TodoRepository interface {
	Store(todo *model.Todo) bool
	FindAll() (model.Todos, error)
}
