package clientaccounts

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAccountsClient,
)
