// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package gitconfig contains configuration utilities for Git.
package gitconfig

import (
	"os/exec"

	"github.com/skeptycal/ghapi/config"
)

type GitConfig interface {
	config.GetSetter
}

func NewGitConfig() GitConfig {
	return &Config{
		username:       "",
		globalPath:     "",
		SessionTimeout: config.DefaultSessionTimeout,
	}
}

func GitPath() (string, error) {
	return exec.LookPath("git")
}

// git config return values
// Reference: https://git-scm.com/docs/git-config
/*
This command will fail with non-zero status upon error. Some exit codes are:
	The section or key is invalid (ret=1),
	no section or name was provided (ret=2),
	the config file is invalid (ret=3),
	the config file cannot be written (ret=4),
	you try to unset an option which does not exist (ret=5),
	you try to unset/set an option for which multiple lines match (ret=5), or
	you try to use an invalid regexp (ret=6).

	On success, the command returns the exit code 0.
*/
