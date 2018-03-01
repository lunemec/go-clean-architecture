package main

import (
	"database/sql"

	"go-clean-architecture/repository/json"
	"go-clean-architecture/repository/mysql"
	"go-clean-architecture/service/robot"
)

func main() {
	// Yes this does not work! Example only!
	db := &sql.DB{}

	importRepo := json.NewImport()
	orgRepo := mysql.NewOrganizationRepository(db)
	locRepo := mysql.NewLocationRepository(db)
	contactRepo := mysql.NewContactRepository(db)

	importService := robot.NewImportService(importRepo, orgRepo, locRepo, contactRepo)
	err := importService.Import(1)
	if err != nil {
		panic(err)
	}
}
