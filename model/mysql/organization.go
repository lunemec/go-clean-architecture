package mysql

import (
	"database/sql"

	"go-clean-architecture/model"
)

// Organization is company standing as IXP association, provider, participant of facility.
type Organization struct {
	ID         uint          `db:"id"`
	ParentID   sql.NullInt64 `db:"parent_id"`
	Name       string
	LocationID sql.NullInt64 `db:"location_id"`
}

// ToModel converts from mysql model to generic model.
func (o *Organization) ToModel() model.Organization {
	return model.Organization{
		// Code omitted.
	}
}

// FromModel converts from generic model to mysql model.
func (o *Organization) FromModel(org model.Organization) Organization {
	return Organization{
		// Code omitted.
	}
}
