package repository

import (
	"database/sql"
	"shiori-server/model"
)

type AccessRepository interface {
	Select() (*model.AccessInfo, error)
}

type accessRepository struct {
	db *sql.DB
}

func NewAccessRepository(db *sql.DB) AccessRepository {
	return &accessRepository{db: db}
}

func (r *accessRepository) Select() (*model.AccessInfo, error) {
	row := r.db.QueryRow(
		"SELECT venue_id, venue_name, venue_address, venue_access_page_url, latitude, longitude, gathering_spot, gathering_date, starting_date, restaurant_name, restaurant_url " +
			"FROM M_ACCESS_INFO " +
			"LIMIT 1",
	)

	ai, err := model.MapToAccessInfo(row)
	if err != nil {
		return nil, err
	}

	return &ai, nil

}
