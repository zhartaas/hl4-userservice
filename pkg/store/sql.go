package store

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLX struct {
	Client *sqlx.DB
}

func New(dsn string, insertExampleData bool) (store SQLX, err error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}

	if err = Migrate(dsn); err != nil {
		err = errors.New("error in migrate:" + err.Error())
	}
	store.Client = db
	fmt.Println("Migrate successfully")

	if insertExampleData {
		err = store.InsertExampleData()
		if err != nil {
			errors.Join(errors.New("Insert example data error: "), err)
		}
	}

	return
}

func (s *SQLX) InsertExampleData() (err error) {
	query := `
    INSERT INTO users (id, name, email, address, role)
    VALUES ('11000000-0000-0000-0000-000000000000', 'example user', 'example.user@gmail.com', 'example address', 'administrator')
    ON CONFLICT (id) DO NOTHING
`

	_, err = s.Client.Exec(query)
	if err != nil {
		err = errors.New("error in insert example data: " + err.Error())
		return
	}

	return
}
