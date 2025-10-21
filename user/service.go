package user

import (
	"go.mod/domain"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
func (svc *service) Get(email string, password string) (*domain.User, error) {

	usr, err := svc.userRepo.Get(email, password)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
