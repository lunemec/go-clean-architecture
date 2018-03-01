package mysql

import (
	"go-clean-architecture/model"
)

type Location struct {
	ID      uint
	Address string
	City    string
	Country string
}

// ToModel converts from mysql model to generic model.
func (l *Location) ToModel() model.Location {
	return model.Location{
		// Code omitted.
	}

}

// FromModel converts from generic model to mysql model.
func (l *Location) FromModel(m model.Location) Location {
	return Location{
		// Code omitted.
	}
}
