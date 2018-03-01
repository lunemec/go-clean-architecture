package json

import (
	"go-clean-architecture/model"
)

type Contact struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// ToModel converts from json model to generic model.
func (l *Contact) ToModel() model.Contact {
	return model.Contact{
		// Code omitted.
	}

}

// FromModel converts from generic model to json model.
func (l *Contact) FromModel(m model.Contact) Contact {
	return Contact{
		// Code omitted.
	}
}
