package connect

import (
	"github.com/aide-family/magicbox/safety"
	"github.com/aide-family/sovereign/pkg/config"
	klog "github.com/go-kratos/kratos/v2/log"
)

var globalRegistry = NewRegistry()

type ORMFactory func(c *config.ORMConfig, logger *klog.Helper) (*ORMConfig, error)

type Registry struct {
	ormConfigs *safety.SyncMap[config.ORMConfig_Dialector, ORMFactory]
}

func (r *Registry) RegisterORMFactory(dialector config.ORMConfig_Dialector, factory ORMFactory) {
	r.ormConfigs.Set(dialector, factory)
}

func (r *Registry) GetORMFactory(dialector config.ORMConfig_Dialector) (ORMFactory, bool) {
	return r.ormConfigs.Get(dialector)
}

func NewRegistry() *Registry {
	return &Registry{
		ormConfigs: safety.NewSyncMap(make(map[config.ORMConfig_Dialector]ORMFactory)),
	}
}
