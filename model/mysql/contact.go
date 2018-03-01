package mysql

import (
	"database/sql"

	"go-clean-architecture/model"
)

// Contact details
type Contact struct {
	ID    uint
	Name  string
	Phone string
	Email string

	OrganizationID sql.NullInt64 `db:"organization_id"`
	LocationID     sql.NullInt64 `db:"location_id"`
}

// ToModel converts from mysql model to generic model.
func (c *Contact) ToModel() model.Contact {
	return model.Contact{
		// Code omitted.
	}
}

// FromModel converts from generic model to mysql model.
func (c *Contact) FromModel(contact model.Contact) Contact {
	return Contact{
		// Code omitted.
	}
}
