package repository

import "github.com/wkmkymt/go-todo-api/domain/model"

// TodoRepository is Todo Repository
type TodoRepository interface {
	FindAll() (model.Users, error)
}
