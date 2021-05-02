package gitconfig

import "os"

type EnvVar interface {
	GetSetter
}

func NewEnvVar(key string) EnvVar {
	e := envVar{key: key}
	_, _ = e.Get()
	return &e
}

type envVar struct {
	key   string
	value string
}

func (e *envVar) String() string {
	// returns only the key
	// prevents accidental return of values from environment variables
	return e.key
}

func (e *envVar) Get() (value string, err error) {
	if e.value == "" {
		s, err := getEnv(e.key)
		if err != nil {
			return "", err
		}
		e.value = s
	}
	return e.value, nil
}

func (e *envVar) Set(value string) (err error) {
	err = os.Setenv(e.key, value)
	if err != nil {
		return err
	}
	e.value = value
	return nil
}

// getEnv returns the value stored in the environment variable key.
//
// Any error returned will be of type *GhError.
//
// If the environment variable is not present, ErrNoEnvironmentVariable
// is returned. If the environment variable is present but the value is
// not set, ErrEnvironmentVariableNotSet is returned. In either case,
// the string value returned with the error is the empty string.
func getEnv(key string) (string, error) {
	token, ok := os.LookupEnv(key)
	if !ok {
		// environment variable is not present
		return "", ErrNoEnvironmentVariable
	}

	// environment variable is present and value is set
	if token != "" {
		return token, nil
	}

	// TODO - check other sources??
	return "", ErrEnvironmentVariableNotSet
}
