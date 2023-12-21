package modulegormsqlite

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
)

const (
	theModuleName     = "github.com/starter-go/module-gorm-sqlite"
	theModuleVersion  = "v1.0.0"
	theModuleRevision = 7
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// NewDriverModule 创建模块 ['github.com/starter-go/module-gorm-sqlite']
func NewDriverModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(libgorm.Module())

	return mb
}
