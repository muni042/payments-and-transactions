//go:build wireinject
// +build wireinject

package main

import (
	"accounts/internal/biz"
	"accounts/internal/conf"
	"accounts/internal/data"
	"accounts/internal/server"
	"accounts/internal/service"
	"accounts/internal/client"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(sc *conf.Server, dc *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		client.ProviderSet,
		newApp,
	))
}
