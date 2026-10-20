//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	clientaccounts "payments/client_accounts"
	clientcard "payments/client_card"
	clienttransactions "payments/client_transactions"
	"payments/internal/biz"
	"payments/internal/conf"
	"payments/internal/data"
	"payments/internal/server"
	"payments/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, clientaccounts.ProviderSet, clienttransactions.ProviderSet, clientcard.ProviderSet, newApp))
}
