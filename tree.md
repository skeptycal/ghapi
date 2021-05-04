# ghapi

> This is the initial directory tree for ghapi. Use the make_tree_md.sh script ([GNU-tree required][get_tree]) to update it if you wish. It is safe to delete this file.

### Directory Structure

```sh
.
├── .VERSION
├── .editorconfig
├── .github
│   ├── FUNDING.yml
│   ├── ISSUE_TEMPLATE
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   ├── dependabot.yml
│   └── workflows
│       ├── codeql-analysis.yml
│       └── go.yml
├── .gitignore
├── .pre-commit-config.yaml
├── CODE_OF_CONDUCT.md
├── LICENSE
├── README.md
├── SECURITY.md
├── apicall.go
├── assets
│   ├── forklist.png
│   └── forks.png
├── cmd
│   ├── apiaccess
│   │   └── main.go
│   ├── context
│   │   └── main.go
│   ├── delrepo
│   │   └── main.go
│   └── example
│       └── ghapi
│           └── main.go
├── config
│   ├── config.go
│   ├── context.go
│   └── errors.go
├── contributing.md
├── coverage.txt
├── docs
│   ├── _config.yml
│   ├── docs.md
│   ├── index.html
│   └── template.md
├── example.go
├── ghapi.go
├── ghconfig
│   ├── example_api.json
│   └── ghconfig.go
├── gitconfig
│   ├── config.go
│   ├── configscopes.go
│   ├── envvar.go
│   ├── error.go
│   ├── getsetter.go
│   ├── gitconfig.go
│   └── gituser.go
├── go.doc
├── go.mod
├── go.sum
├── go.test.sh
├── idea.md
├── make_tree_md.sh
├── notes.md
├── oauth.go
├── permissions
├── repos
├── result.json
├── scripts
│   ├── del_forks.sh
│   ├── forklist
│   ├── forks
│   └── getforks.sh
├── tree.md
└── xhr
    └── xhr.go

16 directories, 58 files
```

[get_tree]: (http://mama.indstate.edu/users/ice/tree/)
