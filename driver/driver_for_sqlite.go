package driver

import (
	"errors"

	"github.com/starter-go/libgorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Driver SQLite 驱动
type Driver struct {

	//starter:component
	_as func(libgorm.Driver) //starter:as(".")

}

func (inst *Driver) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *Driver) Registration() *libgorm.DriverRegistration {
	return &libgorm.DriverRegistration{
		Name:   "sqlite",
		Driver: inst,
	}
}

// Open ...
func (inst *Driver) Open(cfg *libgorm.Configuration) (libgorm.Database, error) {
	db, err := inst.openDB(cfg)
	if err != nil {
		return nil, err
	}
	builder := &libgorm.DatabaseBuilder{DB: db}
	return builder.Create(), nil
}

func (inst *Driver) openDB(cfg *libgorm.Configuration) (*gorm.DB, error) {
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if cfg == nil {
		return nil, errors.New("config==nil")
	}
	path := cfg.Database
	return gorm.Open(sqlite.Open(path), &gorm.Config{})
}
