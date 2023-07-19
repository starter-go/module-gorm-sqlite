package gen4test
import (
    p0ef6f2938 "github.com/starter-go/application"
    p93b03070f "github.com/starter-go/module-gorm-sqlite/src/test/test1"
     "github.com/starter-go/application"
)

// type p93b03070f.TableReg in package:github.com/starter-go/module-gorm-sqlite/src/test/test1
//
// id:com-93b03070f87188ec-test1-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry
// alias:
// scope:singleton
//
type p93b03070f8_test1_TableReg struct {
}

func (inst* p93b03070f8_test1_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-93b03070f87188ec-test1-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p93b03070f8_test1_TableReg) new() any {
    return &p93b03070f.TableReg{}
}

func (inst* p93b03070f8_test1_TableReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p93b03070f.TableReg)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)


    return nil
}


func (inst*p93b03070f8_test1_TableReg) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


