package cli

import (
	"os"

	"github.com/spf13/cobra"
)

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

func flagExists(cmd *cobra.Command, flag string) bool {
	return cmd.Flags().Lookup(flag).Changed
}

// ArgsInCLI checks if there are any arguments in cli
func ArgsInCLI() bool {
	return len(os.Args[1:]) != 0
}
