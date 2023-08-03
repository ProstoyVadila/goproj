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

func (c *Command) Execute() error {
	cmd := exec.Command(c.cmd, c.args...)
	stout, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stout))
	return nil
}
