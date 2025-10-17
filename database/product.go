package database

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

var productList []Product

func Store(pr Product) Product {
	pr.ID = len(productList) + 1
	productList = append(productList, pr)
	return pr
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(pr Product) {
	for idx, product := range productList {
		if product.ID == pr.ID {
			productList[idx] = pr
		}
	}
}

func Delete(productId int) {
	var tmpList []Product
	for _, product := range productList {
		if product.ID != productId {
			tmpList = append(tmpList, product) 
		}
	}
	productList = tmpList
}

func init() {
	productList = append(productList,
		Product{
			ID:          1,
			Title:       "Laptop",
			Description: "A high performance laptop",
			ImageUrl:    "https://example.com/laptop.jpg",
		},
		Product{
			ID:          2,
			Title:       "Smartphone",
			Description: "Latest model smartphone",
			ImageUrl:    "https://example.com/smartphone.jpg",
		},
		Product{
			ID:          3,
			Title:       "Headphones",
			Description: "Noise cancelling headphones",
			ImageUrl:    "https://example.com/headphones.jpg",
		},
	)
}
