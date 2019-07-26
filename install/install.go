package install

import (
	"fmt"

	"github.com/agile-work/srv-mdl-shared/models/job"
	"github.com/agile-work/srv-mdl-shared/models/module"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Init initialize the module installation process
func Init() error {
	mdl := &module.Module{
		ID:     db.UUID(),
		Code:   "task",
		Status: constants.ModuleStatusInstalling,
	}

	trs, err := db.NewTransaction()
	if err != nil {
		return err
	}

	if err := mdl.Register(trs); err != nil {
		trs.Rollback()
		return err
	}

	jobInstance := &job.Instance{}
	id, err := jobInstance.CreateFromJSON(trs, "admin", "./install/module_job_install_v1.0.json", 60, make(map[string]interface{}))
	if err != nil {
		trs.Rollback()
		return err
	}

	fmt.Println("Processing installation through job instance:", id)

	trs.Commit()
	return nil
}
