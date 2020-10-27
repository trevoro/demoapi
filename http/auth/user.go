package auth

import (
	"context"
	"fmt"

	"github.com/trevoro/demoapi/db"
)

const UserContextKey = contextKey("user")

func WithUser(ctx context.Context, user db.User) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

func ContextUser(ctx context.Context) db.User {
	user, ok := ctx.Value(UserContextKey).(db.User)
	// this should never happen but we should still refactor to return an error
	// later anyway.
	if !ok {
		panic(fmt.Errorf("unable to cast context to user"))
	}
	return user
}
