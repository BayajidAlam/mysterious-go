package repo

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageUrl    string    `json:"imageUrl" db:"image_url"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type ProductRepo interface {
	Create(Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(Product) (*Product, error)
}

type productRepo struct {
	db sqlx.DB
}

func NewProductRepo(db sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(pr Product) (*Product, error) {
	query := `
		INSERT INTO products (
			title, 
			description, 
			image_url
		)
		VALUES (
			:title, 
			:description, 
			:image_url
		)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, pr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&pr.ID, &pr.CreatedAt, &pr.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &pr, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	var product Product

	query := `
		SELECT id, 
		title, 
		description, 
		image_url, 
		created_at, 
		updated_at
		FROM products
		WHERE id = $1
		LIMIT 1
	`

	err := r.db.Get(&product, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product

	query := `
		SELECT id, title, description, image_url, created_at, updated_at
		FROM products
		ORDER BY id
	`

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepo) Update(pr Product) (*Product, error) {
	query := `
		UPDATE products
		SET
			title = :title,
			description = :description,
			image_url = :image_url,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = :id
		RETURNING id, title, description, image_url, created_at, updated_at
	`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	var updated Product
	err = stmt.Get(&updated, pr)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *productRepo) Delete(productId int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}

	return nil
}
