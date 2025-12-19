// Package dbimpl is the implementation of the transaction repository for database
package dbimpl

import (
	"context"

	"github.com/aide-family/sovereign/internal/biz/repository"
	"github.com/aide-family/sovereign/internal/data"
	"github.com/aide-family/sovereign/pkg/middler"
	"gorm.io/gorm"
)

func NewTransactionRepository(d *data.Data) repository.Transaction {
	return &transactionRepositoryImpl{
		d: d,
	}
}

type transactionRepositoryImpl struct {
	d *data.Data
}

func (t *transactionRepositoryImpl) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	namespace := middler.GetNamespace(ctx)
	return t.d.BizDB(ctx, namespace).Transaction(func(tx *gorm.DB) error {
		return fn(data.WithBizTransaction(ctx, tx, namespace))
	})
}
