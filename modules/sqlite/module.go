package sqlite

import (
	"github.com/starter-go/application"
	modulegormsqlite "github.com/starter-go/module-gorm-sqlite"
	"github.com/starter-go/module-gorm-sqlite/gen/gen4driver"
)

// Module 函数返回 sqlite 的 libgorm 驱动模块
func Module() application.Module {
	mb := modulegormsqlite.NewDriverModule()
	mb.Components(gen4driver.ExportComForGormSqlite)
	return mb.Create()
}
