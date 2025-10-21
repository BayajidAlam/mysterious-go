package product

import "go.mod/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Get(id int) (*domain.Product, error)
	Delete(pID int) error
	Update(domain.Product) (*domain.Product, error)
}
