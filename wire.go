//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"revil.dev-servers/ggabong"
)

func InitializeGgabong() ggabong.Ggabong {
	wire.Build(ggabong.NewGgabong)

	return ggabong.Ggabong{}
}
