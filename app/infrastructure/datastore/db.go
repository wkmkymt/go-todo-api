package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/wkmkymt/go-todo-api/config"
	"github.com/wkmkymt/go-todo-api/domain/model"
)

// NewDB is Creating DB
func NewDB() *gorm.DB {
	connectString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s %s",
		config.C.Postgres.USER,
		config.C.Postgres.PASS,
		config.C.Postgres.HOST,
		config.C.Postgres.PORT,
		config.C.Postgres.DBNAME,
		config.C.Postgres.OPTION,
	)
	db, err := gorm.Open(config.C.Postgres.DBMS, connectString)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)

	db.AutoMigrate(&model.Todo{})

	return db
}
