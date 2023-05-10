package ggabong

import (
	"fmt"
)

//	m "revil.dev-servers"

type Ggabong struct {
	// Logger *m.Logger
}

/*
	 func NewGgabong(logger *m.Logger) *Ggabong {
		return &Ggabong{
			Logger: logger,
		}
	}
*/
func NewGgabong() Ggabong {
	return Ggabong{}
}

func (g *Ggabong) Start() {
	fmt.Print("test")
}
