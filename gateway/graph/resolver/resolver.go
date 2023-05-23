//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
// "revil.dev-servers/gateway/user"
)

type Resolver struct {
	//userService *user.UserService
}

// userService *user.UserService
func NewResolver() *Resolver {
	return &Resolver{
		//userService: userService,
	}
}
