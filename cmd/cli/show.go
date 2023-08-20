package cli

import (
	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/pkg/output"
	"github.com/spf13/cobra"
)

var showConfigCommand = &cobra.Command{
	Use:     "show",
	Short:   "Demonstrate a global config if it exists",
	Example: "goproj config show",
	Args:    cobra.NoArgs,
	Run:     showConfig,
}

func init() {
	configCommand.AddCommand(showConfigCommand)
}

func showConfig(cmd *cobra.Command, args []string) {
	config, ok := config.Get()
	if !ok {
		output.Info("A Global Config does not set yet")
		return
	}
	config.Show()
}
