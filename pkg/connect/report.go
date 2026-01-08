package connect

import (
	"github.com/aide-family/magicbox/pointer"
	"github.com/aide-family/magicbox/strutil"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kuberegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/aide-family/sovereign/pkg/config"
	"github.com/aide-family/sovereign/pkg/merr"
)

func init() {
	globalRegistry.RegisterReportFactory(config.ReportConfig_KUBERNETES, buildReportFromKubernetes)
	globalRegistry.RegisterReportFactory(config.ReportConfig_ETCD, buildReportFromEtcd)
}

type Report interface {
	registry.Registrar
	registry.Discovery
}

func NewReport(c *config.ReportConfig, logger *klog.Helper) (Report, func() error, error) {
	factory, ok := globalRegistry.GetReportFactory(c.GetReportType())
	if !ok {
		return nil, nil, merr.ErrorInternalServer("report factory not registered")
	}
	return factory(c, logger)
}

func buildReportFromKubernetes(c *config.ReportConfig, logger *klog.Helper) (Report, func() error, error) {
	kubeConfig := &config.KubernetesOptions{}
	if pointer.IsNotNil(c.GetOptions()) {
		if err := anypb.UnmarshalTo(c.GetOptions(), kubeConfig, proto.UnmarshalOptions{Merge: true}); err != nil {
			return nil, nil, merr.ErrorInternalServer("unmarshal kubernetes config failed: %v", err)
		}
	}
	kubeClient, err := NewKubernetesClientSet(kubeConfig.GetKubeConfig())
	if err != nil {
		return nil, nil, merr.ErrorInternalServer("create kubernetes client set failed: %v", err)
	}

	return kuberegistry.NewRegistry(kubeClient, c.GetNamespace()), func() error { return nil }, nil
}

func buildReportFromEtcd(c *config.ReportConfig, logger *klog.Helper) (Report, func() error, error) {
	etcdConfig := &config.ETCDOptions{}
	if pointer.IsNotNil(c.GetOptions()) {
		if err := anypb.UnmarshalTo(c.GetOptions(), etcdConfig, proto.UnmarshalOptions{Merge: true}); err != nil {
			return nil, nil, merr.ErrorInternalServer("unmarshal etcd config failed: %v", err)
		}
	}
	client, err := clientV3.New(clientV3.Config{
		Endpoints:   strutil.SplitSkipEmpty(etcdConfig.GetEndpoints(), ","),
		Username:    etcdConfig.GetUsername(),
		Password:    etcdConfig.GetPassword(),
		DialTimeout: etcdConfig.GetDialTimeout().AsDuration(),
	})
	if err != nil {
		return nil, nil, merr.ErrorInternalServer("create etcd client failed: %v", err)
	}
	return etcd.New(client, etcd.Namespace(c.GetNamespace())), client.Close, nil
}
