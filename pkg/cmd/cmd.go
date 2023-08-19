package cmd

import (
	"fmt"
	"os/exec"
)

type Command struct {
	cmd  string
	args []string
}

func New(cmd string, args ...string) *Command {
	return &Command{
		cmd:  cmd,
		args: args,
	}
}

// commandExists checks existance of the CLI command in the PATH.
func commandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// Execute executes command in command line
func (c *Command) Execute() error {
	if !commandExists(c.cmd) {
		return fmt.Errorf("cannot find this command in the PATH: %s", c.cmd)
	}

	cmd := exec.Command(c.cmd, c.args...)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
