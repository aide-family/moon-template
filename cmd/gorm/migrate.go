package main

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/aide-family/sovereign/cmd"
	"github.com/aide-family/sovereign/internal/biz/do"
)

const cmdMigrateLong = `Migrate database tables based on GORM model definitions`

func newMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database tables based on GORM model definitions",
		Long:  cmdMigrateLong,
		Annotations: map[string]string{
			"group": cmd.DatabaseCommands,
		},
		Run: func(cmd *cobra.Command, args []string) {
			db, closer, err := initDB()
			if err != nil {
				klog.Warnw("msg", "init db failed", "error", err)
				return
			}
			defer closer()
			migrate(db)
		},
	}
}

func migrate(db *gorm.DB) {
	tables := do.Models()
	if err := db.Migrator().AutoMigrate(tables...); err != nil {
		klog.Errorw("msg", "migrate database failed", "error", err, "tables", tables)
		return
	}
	klog.Debugw("msg", "migrate database success")
}
