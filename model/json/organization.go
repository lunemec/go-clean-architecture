package json

import (
	"go-clean-architecture/model"
)

type Organization struct {
	ID       uint          `json:"id"`
	Parent   *Organization `json:"parent"`
	Name     string        `json:"name"`
	Location Location      `json:"location"`

	Contacts []Contact `json:"contacts"`
}

// ToModel converts from json model to generic model.
func (l *Organization) ToModel() model.Organization {
	return model.Organization{
		// Code omitted.
	}

}

// FromModel converts from generic model to json model.
func (l *Organization) FromModel(m model.Organization) Location {
	return Location{
		// Code omitted.
	}
}
