//go:build wireinject
// +build wireinject

// Package grpc is the grpc command for the Sovereign service
package grpc

import (
	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/aide-family/sovereign/cmd/run"
	"github.com/aide-family/sovereign/internal/biz"
	"github.com/aide-family/sovereign/internal/conf"
	"github.com/aide-family/sovereign/internal/data"
	"github.com/aide-family/sovereign/internal/data/impl"
	"github.com/aide-family/sovereign/internal/server"
	"github.com/aide-family/sovereign/internal/service"
)

func WireApp(serviceName string, bc *conf.Bootstrap, helper *klog.Helper) ([]*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSetServerGRPC,
		service.ProviderSetService,
		biz.ProviderSetBiz,
		impl.ProviderSetImpl,
		data.ProviderSetData,
		run.NewApp,
	))
}
