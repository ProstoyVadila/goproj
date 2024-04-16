package git

import (
	"github.com/ProstoyVadila/goproj/pkg/cmd"
)

const (
	gitCommand = "git"
	arg0       = "-C"
	initPath   = "."
	arg2       = "init"
)

// InitGitRepo initialize git repository
func InitGitRepo(path ...string) error {
	arg1 := initPath
	if len(path) != 0 {
		arg1 = path[0]
	}
	command := cmd.New(gitCommand, arg0, arg1, arg2)
	if err := command.Execute(); err != nil {
		return err
	}
	return nil
}
