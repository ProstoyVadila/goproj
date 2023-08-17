package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const FILE = "file"

var configCommand = &cobra.Command{
	Use:     "config",
	Short:   "Set up global configuration for all new generated projects",
	Long:    "Set up global configuration for all new generated projects to not do it every time",
	Example: "goproj config -a \"Bob Doe\" -s=\"Dockerfile,.dockerignore,package/\" --git=false --vscode=false",
	Args:    cobra.NoArgs,
	Run:     setupConfig,
}

func init() {
	rootCommand.AddCommand(configCommand)

	configCommand.Flags().StringP(FILE, "f", "", "an optional flag to set information from yaml file (supprots `json`, `yaml`, `toml`)")
	configCommand.Flags().StringP(AUTHOR, "a", "", "an optional flag to set author name")
	configCommand.Flags().StringSliceP(SKIP, "s", nil, "an optional flag to skip exact files and/or folders (add `/`) from the generation")
	configCommand.Flags().BoolP(GIT, "g", true, "an optional flag to define start git initialization or not")
	configCommand.Flags().BoolP(VSCODE, "c", true, "an optional flag to open the new project in VS Code")
}

func setupConfig(cmd *cobra.Command, args []string) {
	author, err := cmd.Flags().GetString(AUTHOR)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nMY NAME IS %s\n", author)

	config, err := getConfigFile(cmd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
