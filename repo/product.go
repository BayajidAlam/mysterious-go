package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"go.mod/domain"
	"go.mod/product"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db sqlx.DB
}

func NewProductRepo(db sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(pr domain.Product) (*domain.Product, error) {
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

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	var product domain.Product

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

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	var products []*domain.Product

	offset := ((page - 1) * limit) + 1

	query := `
		SELECT id, title, description, image_url, created_at, updated_at
		FROM products
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	err := r.db.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepo) Count() (int64, error) {
	var count int64

	query := `
		SELECT COUNT(*) 
		FROM products
	`

	err := r.db.Get(&count, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *productRepo) Update(pr domain.Product) (*domain.Product, error) {
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

	var updated domain.Product
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
