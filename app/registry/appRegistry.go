package registry

import (
	"github.com/jinzhu/gorm"

	"github.com/wkmkymt/go-todo-api/infrastructure/api/handler"
)

type appInteractor struct {
	todoInteractor todoInteractor
}

// AppInteractor is Application Interactor
type AppInteractor interface {
	NewAppHandler() handler.AppHandler
}

// NewAppInteractor is Creating New Application Interactor
func NewAppInteractor(db *gorm.DB) AppInteractor {
	return &appInteractor{
		todoInteractor{db},
	}
}

func (i *appInteractor) NewAppHandler() handler.AppHandler {
	return handler.AppHandler{
		TodoHandler: i.todoInteractor.NewTodoHandler(),
	}
}
