package product

import "go.mod/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Get(id int) (*domain.Product, error)
	Delete(pID int) error
	Update(domain.Product) (*domain.Product, error)
}
