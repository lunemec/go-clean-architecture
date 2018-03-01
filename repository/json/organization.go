package json

import (
	"go-clean-architecture/model"
)

// Code similar to repository/mysql/Organization but with JSON specific
// code. Get would be HTTP GET, Save HTTP PUT or something similar.

type OrganizationImport struct{}

func NewImport() *OrganizationImport {
	return &OrganizationImport{}
}

func (o *OrganizationImport) Get(id uint) (model.Organization, error) {
	return model.Organization{}, nil
}
