package repository

import (
	"database/sql"
	"shiori-server/model"
)

type GreetingRepository interface {
	FindFirst() (*model.Greeting, error)
}

type greetingRepository struct {
	db *sql.DB
}

func NewGreetingRepository(db *sql.DB) GreetingRepository {
	return &greetingRepository{db: db}
}

func (r *greetingRepository) FindFirst() (*model.Greeting, error) {
	row := r.db.QueryRow(
		"SELECT id, content " +
			"FROM M_GREETING " +
			"WHERE delete_flag = 0 " +
			"ORDER BY id ASC " +
			"LIMIT 1",
	)

	var g model.Greeting
	if err := row.Scan(&g.Id, &g.Content); err != nil {
		return nil, err
	}
	return &g, nil
}
