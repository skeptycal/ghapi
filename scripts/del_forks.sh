#!/usr/bin/env zsh

# reference: https://github.com/cli/cli/issues/2461#issuecomment-732492408
# gh alias set repo-delete 'api -X DELETE "repos/$1"'
# gh repo-delete vilmibm/deleteme

# First, setup the alias:
# gh alias set repo-delete 'api -X DELETE "repos/$1"'
# gh puts configuration files in ~/.config/gh/hosts.yml
# the new alias is the last line in this file

########### file contents
: <<EOF
# What protocol to use when performing git operations. Supported values: ssh, https
git_protocol: https
# What editor gh should run when creating issues, pull requests, etc. If blank, will refer to environment.
editor:
# When to interactively prompt. This is a global config that cannot be overridden by hostname. Supported values: enabled, disabled
prompt: enabled
# A pager program to send command output to, e.g. "less". Set the value to "cat" to disable the pager.
pager:
# Aliases allow you to create nicknames for gh commands
aliases:
    co: pr checkout
    repo-delete: api -X DELETE "repos/$1"
EOF

# set token to normal repo access token using environment variable
GITHUB_TOKEN=$GITHUB_REPO_TOKEN
gh repo list --fork --limit 100 >| forks

# Set token to special 'delete only' token. The only scope is delete repo.
# This way, I can never delete repos by mistake because invoking this token is required.
GITHUB_TOKEN=$GITHUB_DEL_TOKEN

# TODO are there any restrictions on this?
for f in $(cat forklist); do
	if
	gh repo view "$f"
	vared -p 'Delete $f?: ' -c tmp
	# gh repo-delete "$f"
done
GITHUB_TOKEN=$GITHUB_REPO_TOKEN

rm -rf forks


# alternate method:
# reference: https://docs.github.com/en/rest/reference/actions
# curl \
#   -X DELETE \
#   -H "Accept: application/vnd.github.v3+json" \
#   https://api.github.com/repos/octocat/hello-world/actions/artifacts/42
