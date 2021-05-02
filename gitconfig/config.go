package gitconfig

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/skeptycal/ghapi/config"
)

// Config is used to cache configuration values throughout a session to
// reduce response times.
//
// Config caches various properties related to the user, project,
// or server settings. This helps in reducing the response time for every
// request by delivering the properties from a local copy. In majority of
// the cases, the default caching properties should be applicable to your
// infrastructure.
//
// However, there can be scenarios where you may want to modify the
// caching properties.
//
// If you are constantly making schema-level changes to a project that
// affect objects such as attributes, hierarchies, and tables, you may
// want to propagate the changes to users more frequently than the default
// time of 20 minutes.
//
// Modifying the default caching properties is not recommended unless
// you are absolutely sure about the settings you want to change. Using
// the wrong settings can severely affect performance.
type Config struct {
	// username is the configured, authenticated username
	username string

	// currentLevel is the current configuration level that any changes
	// are propagated to.
	currentLevel ConfigScope

	// globalPath is the path for user-level configuration settings.
	globalPath string

	// systemPath is the path for system-level configuration settings.
	systemPath string

	// localPath is the path for local, project-level configuration settings.
	localPath string

	// SessionTimeout - Timeout for the Session Cache, which is specific
	// to a user and exists only while the user is logged in. This cache
	// is deleted when the user logs out or if there’s no activity for a
	// specified timeout interval. The default value for a timeout is 1200
	// seconds or 20 minutes.
	SessionTimeout time.Duration

	// ProjectTimeout - Timeout for the Project Cache, which specific to a
	// project. The default value for a timeout is 86400 seconds or 24 hours.
	ProjectTimeout time.Duration
}

func (c *Config) get(level ConfigScope, key string) (interface{}, error) {
	ctx, cancel := config.GetTimeoutCtx(nil, config.DefaultOsTimeout)
	defer cancel()

	app := "git"
	args := []string{"config", level.Scope(), key}

	cmd := exec.CommandContext(ctx, app, args...)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return nil, ErrEnvironmentVariableNotSet
	}

	return nil, ErrNotImplemented
}

func (c *Config) getCached(level ConfigScope, key string) (string, error) {
	return "", ErrNotImplemented
}

func (c *Config) set(level ConfigScope, key string, value string) error {
	ctx, cancel := config.GetTimeoutCtx(nil, config.DefaultOsTimeout)
	defer cancel()

	app := "git"
	args := []string{"config", level.Scope(), key, value}

	cmd := exec.CommandContext(ctx, app, args...)

	err := cmd.Run()
	if err != nil {
		return ErrEnvironmentVariableNotSet
	}

	return nil
}

func (c *Config) Get(key string) (string, error) {
	v, err := c.get(0, key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), nil
}
func (c *Config) Set(key, value string) error {
	return ErrNotImplemented
}
