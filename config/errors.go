package config

import "errors"

var (
	ErrNotImplemented            = errors.New("feature not yet implemented")
	ErrNoEnvironmentVariable     = errors.New("environment variable not found")
	ErrEnvironmentVariableNotSet = errors.New("environment variable not set")
	ErrNoTokenExists             = errors.New("no valid token exists")
)
