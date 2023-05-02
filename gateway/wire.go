//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeApp() App {
	wire.Build(NewApp)

	return App{}
}
