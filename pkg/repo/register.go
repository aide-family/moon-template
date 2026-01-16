// Package repo is the repository service implementation.
package repo

import (
	"github.com/aide-family/magicbox/safety"

	"github.com/aide-family/sovereign/pkg/config"
	namespacev1 "github.com/aide-family/sovereign/pkg/repo/namespace/v1"
)

var globalRegistry = NewRegistry()

func NewRegistry() Registry {
	return &registry{
		namespaceV1: safety.NewSyncMap(make(map[string]NamespaceFactoryV1)),
	}
}

type NamespaceFactoryV1 func(c *config.NamespaceOptions) (namespacev1.Repository, func() error, error)

type Registry interface {
	RegisterNamespaceV1Factory(name string, factory NamespaceFactoryV1)
	GetNamespaceV1Factory(name string) (NamespaceFactoryV1, bool)
}

type registry struct {
	namespaceV1 *safety.SyncMap[string, NamespaceFactoryV1]
}

func (r *registry) RegisterNamespaceV1Factory(name string, factory NamespaceFactoryV1) {
	r.namespaceV1.Set(name, factory)
}

func (r *registry) GetNamespaceV1Factory(name string) (NamespaceFactoryV1, bool) {
	return r.namespaceV1.Get(name)
}

func RegisterNamespaceV1Factory(name string, factory NamespaceFactoryV1) {
	globalRegistry.RegisterNamespaceV1Factory(name, factory)
}

func GetNamespaceV1Factory(name string) (NamespaceFactoryV1, bool) {
	return globalRegistry.GetNamespaceV1Factory(name)
}
