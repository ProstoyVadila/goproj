package vscode

import (
	"path/filepath"

	"github.com/ProstoyVadila/goproj/pkg/cmd"
)

const (
	initCommand = "code"
	initPath    = "."
)

// InitVSCode opens the new project in VS Code.
func InitVSCode(path ...string) error {
	arg0 := initPath
	if len(path) != 0 {
		arg0 = filepath.Join(path[0], initPath)
	}
	command := cmd.New(initCommand, arg0)
	if err := command.Execute(); err != nil {
		return err
	}
	return nil
}
