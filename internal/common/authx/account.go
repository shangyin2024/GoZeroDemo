package authx

import (
	"context"
)

func GetAccount(ctx context.Context) string {
	userId := ctx.Value("userId")
	uid, ok := userId.(string)
	if !ok {
		panic("userId not found in context")
	}
	return uid
}
