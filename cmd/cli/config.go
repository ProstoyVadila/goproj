package cli

import (
	"log"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/project"
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
	if cmd.Flags().NFlag() == 0 {
		// TODO: there will be a path to input
		log.Println("Specify some flags")
		project.StoreConfig()
		return
	}

	var err error
	conf := new(models.GlobalConfig)

	if flagExists(cmd, FILE) {
		conf, err = reader.GetConfigFile(cmd, FILE)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conf = models.NewGlobalConfig(
			reader.GetAuthor(cmd, AUTHOR),
			reader.GetSkip(cmd, SKIP),
			reader.GetInitGit(cmd, GIT),
			reader.GetVSCode(cmd, VSCODE),
		)
	}

	if err := project.StoreConfig(conf); err != nil {
		log.Fatal(err)
	}
}
