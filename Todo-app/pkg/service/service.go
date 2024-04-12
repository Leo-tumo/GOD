package service

import (
	todo "github.com/Leo-tumo/learngo/Todo-app"
	"github.com/Leo-tumo/learngo/Todo-app/pkg/repository"
	"github.com/fatih/color"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	color.Green("\t\t NEW SERVICE // Initializing...")
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
