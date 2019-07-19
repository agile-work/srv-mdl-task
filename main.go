package main

import (
	shared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-task/install"
	"github.com/agile-work/srv-mdl-task/routes"
)

func main() {
	shared.ListenAndServe("task", "localhost", 3020, installModule, routes.Setup())
}

func installModule() error {
	return install.Init()
}
