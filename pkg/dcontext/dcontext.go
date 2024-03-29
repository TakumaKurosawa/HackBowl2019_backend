package dcontext

import (
	"context"
)

type key string

const (
	userIDKey key = "userID"
)

// SetUserID ContextへUserIDを保存する
func SetUserID(ctx context.Context, authToken string) context.Context {
	return context.WithValue(ctx, userIDKey, authToken)
}

// GetUserIDFromContext ContextからユーザIDを取得する
func GetUserIDFromContext(ctx context.Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID
}
