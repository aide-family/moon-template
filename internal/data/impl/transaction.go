package impl

import (
	"github.com/aide-family/sovereign/internal/biz/repository"
	"github.com/aide-family/sovereign/internal/data"
	"github.com/aide-family/sovereign/internal/data/impl/dbimpl"
)

func NewTransactionRepository(d *data.Data) repository.Transaction {
	return dbimpl.NewTransactionRepository(d)
}
