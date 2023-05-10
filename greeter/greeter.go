package greeter

import (
	"time"

	msg "revil.dev-servers/message"
)

type Greeter struct {
	Grumpy  bool
	Message msg.Message
}

func NewGreeter(m msg.Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() msg.Message {
	if g.Grumpy {
		return msg.Message("Go away!")
	}

	return g.Message
}
