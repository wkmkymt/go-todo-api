package registry

import (
	"github.com/wkmkymt/go-todo-api/infrastructure/api/handler"
	"github.com/wkmkymt/go-todo-api/infrastructure/datastore"
	"github.com/wkmkymt/go-todo-api/interface/controllers"
	"github.com/wkmkymt/go-todo-api/usecase/repository"
	"github.com/wkmkymt/go-todo-api/usecase/service"
)

type todoInteractor struct {}

// TodoInteractor is Todo Interactor
type TodoInteractor interface {
	NewTodoHandler() handler.TodoHandler
}

func (i *todoInteractor) NewTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(i.NewTodoController())
}

func (i *todoInteractor) NewTodoController() controllers.TodoController {
	return controllers.NewTodoController(i.NewTodoService())
}

func (i *todoInteractor) NewTodoService() service.TodoService {
	return service.NewTodoService(i.NewTodoRepository())
}

func (i *todoInteractor) NewTodoRepository() repository.TodoRepository {
	return datastore.NewTodoRepository()
}
