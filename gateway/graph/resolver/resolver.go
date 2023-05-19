//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"revil.dev-servers/gateway/user"
)

type Resolver struct {
	userService *user.UserService
}

func NewResolver(userService *user.UserService) *Resolver {
	return &Resolver{
		userService: userService,
	}
}
