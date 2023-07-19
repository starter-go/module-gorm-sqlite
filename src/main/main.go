package main

import (
	"os"

	modulegormsqlite "github.com/starter-go/module-gorm-sqlite"

	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(modulegormsqlite.Module())
	i.WithPanic(true).Run()
}
