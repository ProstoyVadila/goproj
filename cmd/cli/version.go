package cli

import (
	"github.com/ProstoyVadila/goproj/internal/config/info.go"
	"github.com/spf13/cobra"
)

const VERSION = "version"

var versionCommand = &cobra.Command{
	Use:     "version",
	Short:   "get version info",
	Long:    "get version info",
	Example: "goproj version",
	Args:    cobra.NoArgs,
	Run:     setupVersion,
}

func init() {
	rootCommand.AddCommand(versionCommand)
}

func setupVersion(cmd *cobra.Command, args []string) {
	info.Show()
}
