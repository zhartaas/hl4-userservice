package repository

import (
	"github.com/jmoiron/sqlx"
	"hl4-user_service/userdomain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllUsers() (users []userdomain.Entity, err error) {
	query := `SELECT * FROM users`

	rows, err := r.db.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user userdomain.Entity
		err = rows.StructScan(&user)
		if err != nil {
			return
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}
