package reader

import (
	"errors"
	"log"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/spf13/cobra"
)

// GetPackageName gets a name of package from CLI args.
func GetPackageName(args []string) string {
	if len(args) != 0 {
		return args[0]
	}
	return ""
}

// GetAuthor gets author name of the new project from CLI args.
func GetAuthor(cmd *cobra.Command, flag string) string {
	author, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal(err)
	}
	return author
}

// GetDescription gets description of the new project from CLI args.
func GetDescription(cmd *cobra.Command, flag string) string {
	desc, err := cmd.Flags().GetStringSlice(flag)
	if err != nil {
		log.Fatal(err)
	}
	if len(desc) == 1 {
		return desc[0]
	}
	return ""
}

// GetSkip gets objects to skip in the new project from CLI args.
func GetSkip(cmd *cobra.Command, flag string) []string {
	files, err := cmd.Flags().GetStringSlice(flag)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) != 0 {
		return files
	}
	return make([]string, 0)
}

// GetInitGit gets a boolean value that defines init git repo or not.
func GetInitGit(cmd *cobra.Command, flag string) bool {
	initGit, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatal(err)
	}
	return initGit
}

// GetVSCode gets a boolean balue that defines open the project in VS Code or not.
func GetVSCode(cmd *cobra.Command, flag string) bool {
	initVSCode, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatal(err)
	}
	return initVSCode
}

// GetConfigFile gets information from a configuration file
func GetConfigFile(cmd *cobra.Command, flag string) (config models.GlobalConfig, err error) {
	filename, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal(err)
	}
	if filename == "" {
		return config, errors.New("there is no config file")
	}

	config, err = GetGlobalConfig(filename)
	if err != nil {
		log.Fatal(err)
	}
	return
}
