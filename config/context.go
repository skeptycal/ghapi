package config

import (
	"context"
	"time"
)

const (
	DefaultOsTimeout      = time.Millisecond * 100
	DefaultURLTimeout     = time.Second * 2
	DefaultSessionTimeout = time.Minute * 20
	DefaultProjectTimeout = time.Hour * 24
)

var (
	defaultContext = context.Background()
)

// GetTimeoutCtx returns a context with a timeout.
// If the parent context given is nil, the default context is used.
func GetTimeoutCtx(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if parent == nil {
		parent = defaultContext
	}
	return context.WithTimeout(parent, timeout)
}
