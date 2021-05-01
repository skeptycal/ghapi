# gh api tips

> Handy stuff to know ...

### Deleting repos from the command line

Response by the [trash poet Nate Smith][trash] on [GitHub][issue2461comment732492408]

> As we said in the linked issue:
> We've stayed away from this because it's such a dangerous action and I think we will continue to do that.
> It's possible, however, to generate a PAT with the delete_repo scope and then do:

    gh alias set repo-delete 'api -X DELETE "repos/$1"'
    gh repo-delete vilmibm/deleteme

## Implementing this feature

> The `gh` program is setup to accept one argument

### Setup the alias

To setup the alias, use this one-time command:

```sh
gh alias set repo-delete 'api -X DELETE "repos/$1"'
```

The `gh` program puts configuration files in `~/.config/gh/hosts.yml`. This alias will remain here like any other alias and this step only has to be completed once. The new alias is the last line in this file. The contents should look something like this:

```sh
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
```

### Getting a list of repos

-   we need to grab a list of repos that we want to delete. As noted, I mainly want to get rid of the numerous forks I accumulated over the years. The `gh repo list` command with `--forks` and `--limit` works great. (default --limit is 30) I cache that list in a temp file `forks`.

        gh repo list --fork --limit 100 >| forks

-   If I wanted non-forks (source repos), I would use `--source` instead of `--fork`:

        gh repo list --source --limit 100 >| repos

-   When I go to loop through this list, I want just the first column (the user/repo field):

          awk '{ print $1 }' forks

[trash]: (https://github.com/vilmibm)
[issue2461comment732492408]: (https://github.com/cli/cli/issues/2461#issuecomment-732492408)
