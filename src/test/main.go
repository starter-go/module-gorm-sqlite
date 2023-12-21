package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/module-gorm-sqlite/gen/gen4test"
	"github.com/starter-go/module-gorm-sqlite/modules/sqlite"
	"github.com/starter-go/starter"
)

func main() {

	m := module()

	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "resources"
var theResFS embed.FS

func module() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("module-gorm-sqlite/src/test")
	mb.Version("v1")
	mb.Revision(1)

	mb.EmbedResources(theResFS, "resources")
	mb.Components(gen4test.ExportComForGormSqliteTest)
	mb.Depend(sqlite.Module())

	return mb.Create()
}
