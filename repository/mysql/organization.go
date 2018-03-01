package mysql

import (
	"database/sql"

	"go-clean-architecture/model"
	"go-clean-architecture/model/mysql"
)

// Implements repository/Organization interface.
type Organization struct {
	db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *Organization {
	return &Organization{db: db}
}

func (o *Organization) Get(id uint) (model.Organization, error) {
	// To provide special sql.Null* types or mysql specific ones.
	var org mysql.Organization

	// SELECT FROM organization WHERE id=
	rows, err := o.db.Query("SELECT FROM organization WHERE id=?", id)
	if err != nil {
		return model.Organization{}, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&org); err != nil {
			return model.Organization{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return model.Organization{}, err
	}

	// Return loaded mysql organization as common model without MySQL specific types.
	return org.ToModel(), nil
}

func (o *Organization) Save(model.Organization) (model.Organization, error) {
	// Similar to Get above.
	return model.Organization{}, nil
}
