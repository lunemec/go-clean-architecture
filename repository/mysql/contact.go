package mysql

import (
	"database/sql"

	"go-clean-architecture/model"
)

// Implements repository/Contact interface.
type Contact struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *Contact {
	return &Contact{db: db}
}

func (o *Contact) Get(id uint) (model.Contact, error) {
	// Similar to mysql.Organization.
	return model.Contact{}, nil

}

func (o *Contact) Save(model.Contact) (model.Contact, error) {
	// Similar to mysql.Organization.
	return model.Contact{}, nil
}
