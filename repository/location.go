package repository

import (
	"go-clean-architecture/model"
)

type Location interface {
	Get(uint) (model.Location, error)
	Save(model.Location) (model.Location, error)
}
