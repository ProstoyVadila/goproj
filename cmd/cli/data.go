package cli

import (
	"log"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/reader"
	"github.com/spf13/cobra"
)

// getPackageName gets a name of package from CLI args.
func getPackageName(args []string) string {
	if len(args) != 0 {
		return args[0]
	}
	return ""
}

// getAuthor gets author name of the new project from CLI args.
func getAuthor(cmd *cobra.Command) string {
	author, err := cmd.Flags().GetString(AUTHOR)
	if err != nil {
		log.Fatal(err)
	}
	return author
}

// getDescription gets description of the new project from CLI args.
func getDescription(cmd *cobra.Command) string {
	desc, err := cmd.Flags().GetStringSlice(DESCRIPTION)
	if err != nil {
		log.Fatal(err)
	}
	if len(desc) == 1 {
		return desc[0]
	}
	return ""
}

// getSkip gets objects to skip in the new project from CLI args.
func getSkip(cmd *cobra.Command) []string {
	files, err := cmd.Flags().GetStringSlice(SKIP)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) != 0 {
		return files
	}
	return make([]string, 0)
}

// getFilesToSkip gets files from skip objects.
func getFilesToSkip(cmd *cobra.Command) []string {
	var files []string
	skip := getSkip(cmd)
	for _, object := range skip {
		if object[len(object)-1] != '/' {
			files = append(files, object)
		}
	}
	return files
}

// getFoldersToSkip gets folders from skip objects.
func getFoldersToSkip(cmd *cobra.Command) []string {
	var folders []string
	skip := getSkip(cmd)
	for _, object := range skip {
		last := len(object) - 1
		if object[last] == '/' {
			folders = append(folders, object[:last])
		}
	}
	return folders
}

// getInitGit gets a boolean value that defines init git repo or not.
func getInitGit(cmd *cobra.Command) bool {
	initGit, err := cmd.Flags().GetBool(GIT)
	if err != nil {
		log.Fatal(err)
	}
	return initGit
}

// getVSCode gets a boolean balue that defines open the project in VS Code or not.
func getVSCode(cmd *cobra.Command) bool {
	initVSCode, err := cmd.Flags().GetBool(VSCODE)
	if err != nil {
		log.Fatal(err)
	}
	return initVSCode
}

// getConfigFile gets information from a configuration file
func getConfigFile(cmd *cobra.Command) (config models.ConfigFromFile, err error) {
	filename, err := cmd.Flags().GetString(FILE)
	if err != nil {
		log.Fatal(err)
	}
	if filename == "" {
		return
	}

	config, err = reader.ConfigFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return
}
