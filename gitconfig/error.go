package gitconfig

import (
	"errors"

	"github.com/skeptycal/ghapi/config"
)

var (
	ErrNoTokenExists             = errors.New("no valid token exists")
	ErrNotImplemented            = config.ErrNotImplemented
	ErrNoEnvironmentVariable     = config.ErrNoEnvironmentVariable
	ErrEnvironmentVariableNotSet = config.ErrEnvironmentVariableNotSet
)
