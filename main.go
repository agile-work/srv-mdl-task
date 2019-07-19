package main

import (
	shared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-task/routes"
)

func main() {
	shared.ListenAndServe("task", "localhost", 3020, install, routes.Setup())
}

func install() error {
	return nil
}
