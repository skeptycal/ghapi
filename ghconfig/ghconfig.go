// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package ghconfig contains configuration utilities for Git.
package ghconfig

import "github.com/skeptycal/ghapi/gitconfig"

const (
	defaultGitHubTokenEnvVar = "GITHUB_TOKEN"
	defaultDeleteTokenEnvVar = "GITHUB_DEL_TOKEN"
)

// GetDelToken returns the token stored in the default environment
// variable for the 'delete only' scoped PAT:
//  GITHUB_DEL_TOKEN
//
// If the environment variable is not present or the value is not
// set, the default gh PAT is returned:
//  GITHUB_TOKEN
//
// If neither are present or available, an error is returned.
//
// Any error returned will be of type *GhError.
//
// If the environment variable is not present, ErrNoEnvironmentVariable
// is returned. If the environment variable is present but the value is
// not set, ErrEnvironmentVariableNotSet is returned. In either case,
// the string value returned with the error is the empty string.
func GetDelToken() (string, error) {

	token = gitconfig.NewEnvVar("GITHUB_TOKEN")

	token, err := gitconfig.getEnv(defaultDeleteTokenEnvVar) // TODO - not implemented
	if err == nil {
		return checkToken(token)
	}
	token, err = getEnv(defaultGitHubTokenEnvVar) // TODO - not implemented
	if err == nil {
		return checkToken(token)
	}
	return "", NewGhError("GetDelToken", "GetEnv("+defaultDeleteTokenEnvVar+")", ErrNoEnvironmentVariable)
}

func checkToken(token string) (string, error) {
	return token, nil // TODO - passthrough for now ... until checkToken is finalized
	// return ErrNotImplemented
}
