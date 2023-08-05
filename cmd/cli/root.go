package cli

import "github.com/spf13/cobra"

var rootCommand = &cobra.Command{
	Short: "Go project generator",
	Long:  "Goproj CLI is a tool to generate default template for project in Go",
}

func Execute() error {
	if err := rootCommand.Execute(); err != nil {
		return err
	}
	return nil
}
