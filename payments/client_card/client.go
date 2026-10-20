package clientcard

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCardsClient,
)
