package utils

import (
	"context"
	"time"
)

func InitContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
