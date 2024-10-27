package middlewares

import "context"

// GetUserIDFromContext ...
func GetUserIDFromContext(ctx context.Context) (interface{}, bool) {
	userID, ok := ctx.Value(userIDKey).(interface{})
	return userID, ok
}
