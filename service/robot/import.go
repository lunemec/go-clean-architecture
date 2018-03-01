package robot

import (
	"fmt"

	"go-clean-architecture/repository"
)

// ImportService provides business logic. Suppose it loads something from JSON repository,
// saves in mysql repository.
type ImportService struct {
	jsonRepo    repository.Import
	orgRepo     repository.Organization
	locRepo     repository.Location
	contactRepo repository.Contact
}

// NewImportService is obvious. However the arguments are terrible. Suppose having 50 DB tables, it is hidious.
// But we can't use (repositories ...tableinterface) because we could not be certain all required repositories
// were provided.
// Another issue here is when adding new repositories, the compiler message gets unreadable when missing one of 20 repositories.
func NewImportService(jsonRepo repository.Import, orgRepo repository.Organization, locRepo repository.Location, contactRepo repository.Contact) *ImportService {
	return &ImportService{
		// Code omitted.
	}
}

// Import uses jsonRepository to load data and save them in DB repositories.
func (i *ImportService) Import(id uint) error {
	imported, err := i.jsonRepo.Get(id)
	if err != nil {
		return err
	}

	saved, err := i.orgRepo.Save(imported)
	if err != nil {
		return err
	}

	// Here would be more code handling recursive saving of locations, contacts, etc...
	// Handling cases where you want to update existing records, delete stale etc...

	fmt.Printf("Saved: %v\n", saved)
	return nil
}
