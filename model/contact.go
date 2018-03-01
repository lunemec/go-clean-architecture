package model

type Contact struct {
	ID    uint
	Name  string
	Phone string
	Email string

	Organization *Organization // nil represents no Organization
	Location     *Location     // nil represents no Location
}
