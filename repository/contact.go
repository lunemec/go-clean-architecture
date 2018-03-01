package repository

import (
	"go-clean-architecture/model"
)

type Contact interface {
	Get(uint) (model.Contact, error)
	Save(model.Contact) (model.Contact, error)
}
