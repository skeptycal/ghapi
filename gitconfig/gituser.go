package gitconfig

type GitUser interface {
	config.User
}

func NewGitUser() GitUser {
	return &gitUser{
		username:   "",
		ghusername: "",
	}
}

type gitUser struct {
	username   string
	ghusername string
}

func (c *gitUser) SetUserName(name string) error {
	// git config --global user.name
	return ErrNotImplemented
}

func (c *gitUser) GetUserName() (string, error) {
	// git config --global user.name
	return "", ErrNotImplemented

}
