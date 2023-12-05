package test1

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
)

const (
	theNamespace = "module-gorm-sqlite/src/test/test1"
)

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.GroupRegistry) //starter:as(".")

	Context application.Context //starter:inject("context")

	Enabled string //starter:inject("${datagroup.dg1.enabled}")
	Prefix  string //starter:inject("${datagroup.dg1.table-name-prefix}")
	URI     string //starter:inject("${datagroup.dg1.uri}")
	Source  string //starter:inject("${datagroup.dg1.datasource}")

	// agent  libgorm.DatabaseAgent
	agent libgorm.Agent
}

func (inst *TableReg) _impl() libgorm.GroupRegistry {
	return inst
}

// Groups ...
func (inst *TableReg) Groups() []*libgorm.GroupRegistration {

	r1 := &libgorm.GroupRegistration{
		Enabled: true,
		Alias:   "",
		Prefix:  inst.Prefix,
		URI:     inst.URI,
		Source:  inst.Source,
		Group:   inst,
	}

	return []*libgorm.GroupRegistration{r1}
}

// Prototypes 列出各种 entity 的原型
func (inst *TableReg) Prototypes() []any {
	list := make([]any, 0)
	list = append(list, &Table1{})
	list = append(list, &Table2{})
	list = append(list, &Table3{})
	return list
}
