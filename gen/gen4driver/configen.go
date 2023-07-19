package gen4driver

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForGormSqlite ...
func ExportComForGormSqlite(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
