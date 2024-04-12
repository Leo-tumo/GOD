package repository

import (
	todo "github.com/Leo-tumo/learngo/Todo-app"
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	color.Cyan("\t\tNEW REPOSITORY")
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
