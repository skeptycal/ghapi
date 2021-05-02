// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package config contains configuration utilities for cli applications.
package config

import (
	"context"
	"time"
)

type User interface {
	GetSetter
}

type Getter interface {
	Get() (value string, err error)
}

type Setter interface {
	Set(value string) (err error)
}

type GetSetter interface {
	Getter
	Setter
}

type Contexter interface {
	WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
}
