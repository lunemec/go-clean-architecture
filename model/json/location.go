package json

import (
	"go-clean-architecture/model"
)

type Location struct {
	ID      uint   `json:"id"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// ToModel converts from json model to generic model.
func (l *Location) ToModel() model.Location {
	return model.Location{
		// Code omitted.
	}

}

// FromModel converts from generic model to json model.
func (l *Location) FromModel(m model.Location) Location {
	return Location{
		// Code omitted.
	}
}
