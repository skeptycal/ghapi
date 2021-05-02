package main

// delrepo is a command line utility to delete a github repo.
// it requires a developer 'delete only' PAT setup prior to use.

import (
	"fmt"
	"os"

	"github.com/skeptycal/ghapi"
)

func run() error {
	defaultGhDelToken := "GITHUB_DEL_TOKEN"

	delToken, err := ghapi.GetEnv(defaultGhDelToken)
	if err != nil {
		return err
	}
	oldToken, err := ghapi.GetEnv(defaultGhDelToken)
	if err != nil {
		return err
	}
	defer os.Setenv("GITHUB_TOKEN", oldToken)
	defer os.Setenv("GITHUB_TOKEN", "")

	os.Setenv("GITHUB_TOKEN", delToken)

	fmt.Println("GitHub 'delete only' scope token: ", delToken)
	return nil
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Fprintf(os.Stderr, "Fatal panic: %v", r)
			os.Exit(1)
		}
	}()

	err := run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in execution: %v\n", err)
		os.Exit(1)
	}
}
