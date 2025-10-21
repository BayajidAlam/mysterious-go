package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"go.mod/domain"
	"go.mod/user"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
        INSERT INTO users (
				first_name, 
				last_name, 
				email, 
				password, 
				is_shop_owner
				)
        VALUES (
				:first_name, 
				:last_name, 
				:email, 
				:password, 
				:is_shop_owner
				)
        RETURNING 
				id
    `

	var userId int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
	}

	user.ID = userId
	return &user, nil
}

func (r *userRepo) Get(email string, password string) (*domain.User, error) {
	var user domain.User
	query := `
        SELECT id, 
				first_name, 
				last_name, 
				email, 
				password, 
				is_shop_owner
        FROM users
        WHERE email = $1 AND password = $2
        LIMIT 1
    `
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
