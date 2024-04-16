package vscode

import (
	"path/filepath"

	"github.com/ProstoyVadila/goproj/pkg/cmd"
)

const (
	initCommand = "code"
	arg0        = "."
)

// InitVSCode opens the new project in VS Code.
func InitVSCode(path ...string) error {
	arg := arg0
	if len(path) != 0 {
		arg = filepath.Join(path[0], arg0)
	}
	command := cmd.New(initCommand, arg)
	if err := command.Execute(); err != nil {
		return err
	}
	return nil
}
