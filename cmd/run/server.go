// Package run is the run command for the Sovereign service
package run

import (
	"github.com/go-kratos/kratos/v2/config/env"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"

	"github.com/aide-family/sovereign/internal/conf"
)

const cmdRunLong = `Run the Sovereign services`

func NewCmd(defaultServerConfigBytes []byte) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the Sovereign services",
		Long:  cmdRunLong,
	}
	var bc conf.Bootstrap
	if err := conf.Load(&bc, env.NewSource(), conf.NewBytesSource(defaultServerConfigBytes)); err != nil {
		klog.Errorw("msg", "load config failed", "error", err)
		panic(err)
	}
	runFlags.addFlags(runCmd, &bc)

	return runCmd
}
