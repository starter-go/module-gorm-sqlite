package gen4test
import (
    p0ef6f2938 "github.com/starter-go/application"
    p93b03070f "github.com/starter-go/module-gorm-sqlite/src/test/test1"
     "github.com/starter-go/application"
)

// type p93b03070f.TableReg in package:github.com/starter-go/module-gorm-sqlite/src/test/test1
//
// id:com-93b03070f87188ec-test1-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p93b03070f8_test1_TableReg struct {
}

func (inst* p93b03070f8_test1_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-93b03070f87188ec-test1-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
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
    com.Enabled = inst.getEnabled(ie)
    com.Prefix = inst.getPrefix(ie)
    com.URI = inst.getURI(ie)
    com.Source = inst.getSource(ie)


    return nil
}


func (inst*p93b03070f8_test1_TableReg) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p93b03070f8_test1_TableReg) getEnabled(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.dg1.enabled}")
}


func (inst*p93b03070f8_test1_TableReg) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.dg1.table-name-prefix}")
}


func (inst*p93b03070f8_test1_TableReg) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.dg1.uri}")
}


func (inst*p93b03070f8_test1_TableReg) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.dg1.datasource}")
}


