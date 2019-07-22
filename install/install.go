package install

import (
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

	jobInstance := &job.Instance{}
	if err := jobInstance.CreateFromJSON(trs, "admin", "./install/test.json", 60, make(map[string]interface{})); err != nil {
		trs.Rollback()
		return err
	}

	if err := mdl.Register(trs); err != nil {
		trs.Rollback()
		return err
	}

	trs.Commit()
	return nil
}
