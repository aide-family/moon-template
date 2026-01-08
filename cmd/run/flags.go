package run

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aide-family/magicbox/load"
	"github.com/aide-family/magicbox/pointer"
	"github.com/aide-family/magicbox/strutil"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/aide-family/sovereign/cmd"
	"github.com/aide-family/sovereign/internal/conf"
	"github.com/aide-family/sovereign/pkg/enum"
)

type RunFlags struct {
	*conf.Bootstrap
	*cmd.GlobalFlags

	metadata    []string
	useRandomID bool
	configPaths []string
	environment string
	jwtExpire   string
}

var runFlags RunFlags

func (f *RunFlags) addFlags(c *cobra.Command, bc *conf.Bootstrap) {
	f.GlobalFlags = cmd.GetGlobalFlags()
	f.Bootstrap = bc

	c.PersistentFlags().StringSliceVarP(&f.configPaths, "config", "c", []string{}, `Example: -c=./config1/ -c=./config2/`)

	c.PersistentFlags().StringVar(&f.Server.Name, "server-name", f.Server.Name, `Example: --server-name="sovereign"`)
	useRandomID, _ := strconv.ParseBool(f.Server.UseRandomID)
	c.PersistentFlags().BoolVar(&f.useRandomID, "use-random-node-id", useRandomID, `Example: --use-random-node-id`)
	metadataStr := make([]string, 0, len(f.Server.Metadata))
	for key, value := range f.Server.Metadata {
		metadataStr = append(metadataStr, fmt.Sprintf("%s=%s", key, value))
	}
	c.PersistentFlags().StringSliceVar(&f.metadata, "server-metadata", metadataStr, `Example: --server-metadata="tag=sovereign" --server-metadata="email=aidecloud@163.com"`)
	c.PersistentFlags().StringVar(&f.environment, "environment", f.Environment.String(), `Example: --environment="DEV", --environment="TEST", --environment="PREVIEW", --environment="PROD"`)
	c.PersistentFlags().StringVar(&f.Jwt.Secret, "jwt-secret", f.Jwt.Secret, `Example: --jwt-secret="xxx"`)
	c.PersistentFlags().StringVar(&f.jwtExpire, "jwt-expire", f.Jwt.Expire.AsDuration().String(), `Example: --jwt-expire="10s", --jwt-expire="1m", --jwt-expire="1h", --jwt-expire="1d"`)
	c.PersistentFlags().StringVar(&f.Jwt.Issuer, "jwt-issuer", f.Jwt.Issuer, `Example: --jwt-issuer="xxx"`)
}

func (f *RunFlags) ApplyToBootstrap() error {
	if strutil.IsEmpty(f.Server.Name) {
		f.Server.Name = f.Name
	}
	if strutil.IsEmpty(f.Server.Namespace) {
		f.Server.Namespace = f.Namespace
	}

	metadata := f.Server.Metadata
	if pointer.IsNil(metadata) {
		metadata = make(map[string]string)
	}

	metadata["repository"] = f.Repo
	metadata["author"] = f.Author
	metadata["email"] = f.Email
	metadata["built"] = f.Built

	for _, m := range f.metadata {
		parts := strings.SplitN(m, "=", 2)
		if len(parts) == 2 {
			metadata[parts[0]] = parts[1]
		}
	}

	f.Server.Metadata = metadata
	f.Server.UseRandomID = strconv.FormatBool(f.useRandomID)

	if strutil.IsNotEmpty(f.environment) {
		f.Environment = enum.Environment(enum.Environment_value[f.environment])
	}

	if strutil.IsNotEmpty(f.jwtExpire) {
		if expire, err := time.ParseDuration(f.jwtExpire); pointer.IsNil(err) {
			f.Jwt.Expire = durationpb.New(expire)
		}
	}

	if len(f.configPaths) > 0 {
		var bc conf.Bootstrap
		sourceOpts := make([]kconfig.Source, 0, len(f.configPaths))
		sourceOpts = append(sourceOpts, env.NewSource())
		for _, configPath := range f.configPaths {
			if strutil.IsNotEmpty(configPath) {
				sourceOpts = append(sourceOpts, file.NewSource(load.ExpandHomeDir(strings.TrimSpace(configPath))))
			}
		}
		if len(sourceOpts) > 0 {
			if err := conf.Load(&bc, sourceOpts...); err != nil {
				klog.Warnw("msg", "load config failed", "error", err)
				return err
			}
			f.Bootstrap = &bc
		}
	}

	return nil
}

func GetRunFlags() *RunFlags {
	return &runFlags
}
