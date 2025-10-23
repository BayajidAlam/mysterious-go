package product

import (
	"go.mod/domain"
	productHandler "go.mod/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Delete(productId int) error
	Update(domain.Product) (*domain.Product, error)
}

func (svc *service) Create(prdct domain.Product) (*domain.Product, error) {
	return svc.productRepo.Create(prdct)
}
func (svc *service) Get(productID int) (*domain.Product, error) {
	return svc.productRepo.Get(productID)
}
func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.productRepo.List(page, limit)
}
func (svc *service) Count() (int64, error) {
	return svc.productRepo.Count()
}
func (svc *service) Delete(productId int) error {
	return svc.productRepo.Delete(productId)
}
func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.productRepo.Update(prdct)
}
