package utils

import "context"

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	UserClaimsContextKey = contextKey("userClaims")
)

func GetUserClaimsFromContext(ctx context.Context) (string, bool) {
	caller, ok := ctx.Value(UserClaimsContextKey).(string)
	return caller, ok
}
