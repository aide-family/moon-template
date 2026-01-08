// Package gorm is the gorm package for the Rabbit service
package main

import (
	"github.com/aide-family/magicbox/log"
	"github.com/aide-family/magicbox/log/stdio"
	"github.com/aide-family/magicbox/strutil"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/aide-family/sovereign/cmd"
	"github.com/aide-family/sovereign/internal/conf"
	"github.com/aide-family/sovereign/pkg/connect"
	"github.com/aide-family/sovereign/pkg/merr"
)

const cmdLong = `GORM code generation and database migration tools for Sovereign service`

func init() {
	logger, err := log.NewLogger(stdio.LoggerDriver())
	if err != nil {
		panic(merr.ErrorInternal("new logger failed with error: %v", err).WithCause(err))
	}
	logger = klog.With(logger,
		"ts", klog.DefaultTimestamp,
	)
	filterLogger := klog.NewFilter(logger, klog.FilterLevel(klog.LevelInfo))
	helper := klog.NewHelper(filterLogger)
	klog.SetLogger(helper.Logger())
}

func NewCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "gorm",
		Short: "GORM code generation and database migration tools",
		Long:  cmdLong,
		Annotations: map[string]string{
			"group": cmd.ServiceCommands,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	flags.addFlags(runCmd)
	runCmd.AddCommand(
		newGenCmd(),
		newMigrateCmd(),
	)
	return runCmd
}

func initDB() (*gorm.DB, func() error, error) {
	flags.GlobalFlags = cmd.GetGlobalFlags()

	var bc conf.Bootstrap
	if strutil.IsNotEmpty(flags.configPath) {
		klog.Debugw("msg", "load config file", "file", flags.configPath)
		c := config.New(config.WithSource(
			env.NewSource(),
			file.NewSource(flags.configPath),
		))
		if err := c.Load(); err != nil {
			klog.Errorw("msg", "load config failed", "error", err)
			return nil, nil, err
		}

		if err := c.Scan(&bc); err != nil {
			klog.Errorw("msg", "scan config failed", "error", err)
			return nil, nil, err
		}
	}
	flags.applyToBootstrap(&bc)

	db, closer, err := connect.NewDB(flags.ormConfig, klog.NewHelper(klog.GetLogger()))
	if err != nil {
		klog.Errorw("msg", "new db failed", "error", err)
		return nil, nil, err
	}
	return db.Debug(), closer, nil
}

func main() {
	rootCmd := cmd.NewCmd()
	cmd.Execute(rootCmd, NewCmd())
}
