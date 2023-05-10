//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeApp() (App, error) {
	wire.Build(NewApp, NewConfig, NewServer)

	return App{}, nil
}
