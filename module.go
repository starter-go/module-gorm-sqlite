package modulegormsqlite

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modgorm"
	"github.com/starter-go/module-gorm-sqlite/gen/gen4driver"
)

const (
	theModuleName     = "github.com/starter-go/module-gorm-sqlite"
	theModuleVersion  = "v0.9.0"
	theModuleRevision = 3
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块 ['github.com/starter-go/module-gorm-sqlite']
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4driver.ExportComForGormSqlite)

	mb.Depend(modgorm.Module())

	return mb.Create()
}
