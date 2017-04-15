package utility

import (
	"fmt"

	"golang.org/x/net/context"
)

type userIDKey string

const ContextUserIDKey userIDKey = "tokenKey"

func SetUserID(parents context.Context, userID int) context.Context {
	return context.WithValue(parents, ContextUserIDKey, userID)
}

func GetUserID(ctx context.Context) (int, error) {
	v := ctx.Value(ContextUserIDKey)

	userID, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("token not found")
	}

	return userID, nil
}

type tokenKey string

const ContextTokenKey userIDKey = "tokenKey"

func SetToken(parents context.Context, userID int) context.Context {
	return context.WithValue(parents, ContextTokenKey, userID)
}

func GetToken(ctx context.Context) (string, error) {
	v := ctx.Value(ContextTokenKey)

	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}
