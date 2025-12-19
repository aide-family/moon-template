// Package biz is the business logic for the Sovereign service.
package biz

import "github.com/google/wire"

var ProviderSetBiz = wire.NewSet(
	NewHealth,
	NewNamespace,
)
