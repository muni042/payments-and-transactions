package clienttransactions

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTransactionsClient,
)
