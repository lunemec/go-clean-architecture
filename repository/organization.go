package repository

import (
	"go-clean-architecture/model"
)

type Organization interface {
	Get(uint) (model.Organization, error)
	Save(model.Organization) (model.Organization, error)
}

// Specific interface for JSON import.
type Import interface {
	Get(uint) (model.Organization, error)
}
