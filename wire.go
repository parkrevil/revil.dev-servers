//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"revil.dev-servers/event"
	"revil.dev-servers/greeter"
	"revil.dev-servers/message"
)

func InitializeEvent(phrase string) (event.Event, error) {
	wire.Build(event.NewEvent, greeter.NewGreeter, message.NewMessage)

	return event.Event{}, nil
}
