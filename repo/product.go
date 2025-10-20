package repo

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type ProductRepo interface {
	Create(Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	GenerateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(pr Product) (*Product, error) {
	pr.ID = len(r.productList) + 1
	r.productList = append(r.productList, &pr)
	return &pr, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productId int) error {
	var tmpList []*Product
	for _, product := range r.productList {
		if product.ID != productId {
			tmpList = append(tmpList, product)
		}
	}
	r.productList = tmpList

	return nil
}

func (r *productRepo) Update(pr Product) (*Product, error) {
	for idx, product := range r.productList {
		if product.ID == pr.ID {
			r.productList[idx] = &pr
		}
	}

	return &pr, nil
}

func GenerateInitialProducts(r *productRepo) {
	r.productList = append(r.productList,
		&Product{
			ID:          1,
			Title:       "Laptop",
			Description: "A high performance laptop",
			ImageUrl:    "https://example.com/laptop.jpg",
		},
		&Product{
			ID:          2,
			Title:       "Smartphone",
			Description: "Latest model smartphone",
			ImageUrl:    "https://example.com/smartphone.jpg",
		},
		&Product{
			ID:          3,
			Title:       "Headphones",
			Description: "Noise cancelling headphones",
			ImageUrl:    "https://example.com/headphones.jpg",
		},
	)
}
