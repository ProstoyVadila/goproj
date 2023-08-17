package vscode

import (
	"github.com/ProstoyVadila/goproj/pkg/cmd"
)

const (
	initCommand = "code"
	arg0        = "."
)

// InitVSCode opens the new project in the VS Code.
func InitVSCode() error {
	command := cmd.New(initCommand, arg0)
	if err := command.Execute(); err != nil {
		return err
	}
	return nil
}
