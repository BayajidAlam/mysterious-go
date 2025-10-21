package user

import "go.mod/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Get(email string, password string) (*domain.User, error)
}
