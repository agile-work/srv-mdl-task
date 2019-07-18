package main

import (
	shared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-task/routes"
)

func main() {
	shared.ListenAndServe("Task", "localhost", 3020, routes.Setup())
}
