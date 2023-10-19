package auth

import (
	"context"
	"errors"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	isAdmin bool
}

func (u User) IsAdmin() bool {
	return u.isAdmin
}

var (
	ErrUserNotFound = errors.New("user not found")

	loginFn = inMemLogIn

	ctxUserKey = struct{}{}
)

func LogIn(ctx context.Context, login, password string) (context.Context, error) {
	user, err := loginFn(ctx, login, password)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, ctxUserKey, user), nil
}

func RequireAdmin(ctx context.Context) bool {
	user, ok := ctx.Value(ctxUserKey).(User)
	return ok && user.isAdmin
}

func CtxUser(ctx context.Context) (u User, ok bool) {
	u, ok = ctx.Value(ctxUserKey).(User)
	return
}
