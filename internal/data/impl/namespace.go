package impl

import (
	"github.com/aide-family/sovereign/internal/biz/repository"
	"github.com/aide-family/sovereign/internal/data"
	"github.com/aide-family/sovereign/internal/data/impl/dbimpl"
)

func NewNamespaceRepository(d *data.Data) repository.Namespace {
	return dbimpl.NewNamespaceRepository(d)
}
