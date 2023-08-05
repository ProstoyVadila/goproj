package cli

import (
	"fmt"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/project"
	"github.com/spf13/cobra"
)

const (
	SKIP        = "skip"
	AUTHOR      = "author"
	DESCRIPTION = "description"
)

var packageCommand = &cobra.Command{
	Use:   "init",
	Short: "Generate a new Go porject with default files and folders",
	Long:  "Generate a new Go porject with default files (README.md, LICENSE, go.mod, Makefile, Dockerfile, .gitignore, .dockerignore, .env) and folders (cmd/, internal/, pkg/, tests/)",
	Run:   packageName,
}

func init() {
	rootCommand.AddCommand(packageCommand)

	packageCommand.PersistentFlags().StringP(AUTHOR, "a", "", "an optional flag to set your name")
	packageCommand.PersistentFlags().StringSliceP(DESCRIPTION, "d", nil, "an optional flag to set a description of your project")
	packageCommand.PersistentFlags().StringSliceP(SKIP, "e", nil, "an optional flag to exclude exact files from the generation")
}

// packageName gets poject's setup from CLI and runs generation and initialization of files/git repo.
func packageName(cmd *cobra.Command, args []string) {

	setup := models.NewSetup(
		getPackageName(args),
		getAuthor(cmd),
		getDescription(cmd),
		getFilesToSkip(cmd),
	)

	showSetup(setup)
	project.Generate(setup)
}

// showSetup writes setup info from CLI to standart output.
func showSetup(setup *models.Setup) {
	fmt.Printf("\nProject (package) name: %s\n", setup.PackageName)
	fmt.Printf("Author: %s\n", setup.Author)
	fmt.Printf("Description: %s\n", setup.Description)
	fmt.Printf("Files to skip: %v\n\n", setup.FilesToSkip)
}
