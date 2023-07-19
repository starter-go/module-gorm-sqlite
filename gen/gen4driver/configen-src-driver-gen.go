package gen4driver
import (
    p37720cfba "github.com/starter-go/module-gorm-sqlite/driver"
     "github.com/starter-go/application"
)

// type p37720cfba.Driver in package:github.com/starter-go/module-gorm-sqlite/driver
//
// id:com-37720cfbad7ec1d3-driver-Driver
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-Driver
// alias:
// scope:singleton
//
type p37720cfbad_driver_Driver struct {
}

func (inst* p37720cfbad_driver_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-37720cfbad7ec1d3-driver-Driver"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p37720cfbad_driver_Driver) new() any {
    return &p37720cfba.Driver{}
}

func (inst* p37720cfbad_driver_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p37720cfba.Driver)
	nop(ie, com)

	


    return nil
}


