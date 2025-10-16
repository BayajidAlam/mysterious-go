package database

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

var ProductList []Product

func init() {
	ProductList = append(ProductList,
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
