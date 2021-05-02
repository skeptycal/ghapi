package gitconfig

// ConfigScope is the configuration scope used to specify the
// configuration file location
//
// Config file location
//
// --global              use global config file (default)
//
// --system              use system config file
//
// --local               use repository config file
//
// The following options are not implemented:
//
// --worktree            use per-worktree config file
//
// -f, --file <file>     use given config file
//
// --blob <blob-id>      read config from given blob object
//
// Reference: https://git-scm.com/docs/git-config#FILES
type ConfigScope int

// ConfigScope is the configuration scope used to specify the
// configuration file location
//
// Config file location
//
// --global              use global config file (default)
//
// --system              use system config file
//
// --local               use repository config file
//
// The following options are not implemented:
//
// --worktree            use per-worktree config file
//
// -f, --file <file>     use given config file
//
// --blob <blob-id>      read config from given blob object
//
// Reference: https://git-scm.com/docs/git-config#FILES

const (
	Global ConfigScope = iota
	System
	Local
	// worktree		// Not implemented
	// file			// Not implemented
	// blob			// Not implemented
)

func (c ConfigScope) String() string {
	return configScopes[int(c)].name
}

func (c ConfigScope) Scope() string {
	return "--" + c.String()
}

func (c ConfigScope) FilePath() string {
	return configScopes[int(c)].filepath
}

// func updateConfigScopes(global, system, local string) {
// 	configScopes[0] = {"global",global}
// 	configScopes[1] = {"system",system}
// 	configScopes[2] = {"local",local}
// }

var configScopes = map[int]struct{ name, filepath string }{
	0: {"global", ""},
	1: {"system", ""},
	2: {"local", ""},
}
