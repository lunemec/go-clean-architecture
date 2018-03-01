package mysql

import (
	"database/sql"

	"go-clean-architecture/model"
)

// Implements repository/Location interface.
type Location struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *Location {
	return &Location{db: db}
}

func (o *Location) Get(id uint) (model.Location, error) {
	// Similar to mysql.Organization.
	return model.Location{}, nil
}

func (o *Location) Save(model.Location) (model.Location, error) {
	// Similar to mysql.Organization.
	return model.Location{}, nil
}
