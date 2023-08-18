package cli

import (
	"fmt"
	"log"

	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/reader"
	"github.com/spf13/cobra"
)

const FILE = "file"

var configCommand = &cobra.Command{
	Use:     "config",
	Short:   "Set up global configuration for all new generated projects",
	Long:    "Set up global configuration for all new generated projects to not do it every time. You can find generated config file \".goproj.config.toml\" in your user folder and change it manually.",
	Example: "goproj config -a \"Bobert Doe\" -s=\"Dockerfile,.dockerignore,internal/,pkg/\" -g=false --vscode=false",
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
	conf, err := reader.GetConfigFile(cmd, FILE)
	if err != nil {
		conf.Author = reader.GetAuthor(cmd, AUTHOR)
		conf.Description = reader.GetDescription(cmd, DESCRIPTION)
		conf.InitGit = reader.GetInitGit(cmd, GIT)
		conf.InitVSCode = reader.GetVSCode(cmd, VSCODE)
		conf.Skip = reader.GetSkip(cmd, SKIP)
	}

	err = config.Store(conf)
	if err != nil {
		log.Fatal(err)
	}

	showConfig(conf)
}

func showConfig(conf models.ConfigFromFile) {
	fmt.Println("\nYou set this configuration for your new projects:")
	fmt.Printf("Author: %s\n", conf.Author)
	fmt.Printf("Description: %s\n", conf.Description)
	fmt.Printf("Obejcts to Skip: %v\n\n", conf.Skip)
}
