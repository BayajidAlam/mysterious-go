package user

import (
	"go.mod/domain"
	userHandler "go.mod/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Get(email string, password string) (*domain.User, error)
}
