package model

type Organization struct {
	ID       uint
	Parent   *Organization
	Name     string
	Location Location

	Contacts []Contact
}
