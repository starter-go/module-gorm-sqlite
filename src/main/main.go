package main

import (
	"os"

	"github.com/starter-go/module-gorm-sqlite/modules/sqlite"

	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(sqlite.Module())
	i.WithPanic(true).Run()
}
