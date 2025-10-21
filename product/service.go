package product

type service struct {
	productRepo ProductRepo
}

func NewService(ProductRepo ProductRepo) *service {
	return &service{
		productRepo: ProductRepo,
	}
}
