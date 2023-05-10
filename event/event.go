package event

import (
	"errors"
	"fmt"

	greeter "revil.dev-servers/greeter"
)

type Event struct {
	Greeter greeter.Greeter
}

func NewEvent(g greeter.Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}

	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()

	fmt.Println(msg)
}
