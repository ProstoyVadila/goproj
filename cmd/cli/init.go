package cli

import (
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/project"
	"github.com/ProstoyVadila/goproj/pkg/reader"
	"github.com/spf13/cobra"
)

var initPackageCommand = &cobra.Command{
	Use:     "init",
	Short:   "Generate a new Go porject with default files and folders",
	Long:    "Generate a new Go porject with default files (README.md, LICENSE, go.mod, Makefile, Dockerfile, .gitignore, .dockerignore, .env) and folders (cmd/, internal/, pkg/, tests/)",
	Example: "goproj new github.com/Bobert/new_project -a \"Bobert Doe\" -d=\"My new project\" -s=\"Dockerfile,.dockerignore,internal/,pkg/\" --git=false",
	Args:    cobra.ExactArgs(1),
	Run:     initPackageName,
}

func init() {
	rootCommand.AddCommand(initPackageCommand)

	initPackageCommand.Flags().StringP(AUTHOR, "a", "", "an optional flag to set your name")
	initPackageCommand.Flags().StringSliceP(DESCRIPTION, "d", nil, "an optional flag to set a description of your project")
	initPackageCommand.Flags().StringSliceP(SKIP, "s", nil, "an optional flag to skip exact files and/or folders (add `/` after folder's name) from the generation.")
	initPackageCommand.Flags().BoolP(GIT, "g", false, "an optional flag to define start git initialization or not (Default false)")
	initPackageCommand.Flags().BoolP(VSCODE, "c", false, "an optional flag to open the new project in VS Code (Default false)")
}

// initPackageName gets poject's setup from CLI and runs generation and initialization of files/git repo.
func initPackageName(cmd *cobra.Command, args []string) {
	const generateNewFolder = false
	setup := models.NewSetup(
		reader.GetPackageName(args),
		reader.GetAuthor(cmd, AUTHOR),
		reader.GetDescription(cmd, DESCRIPTION),
		reader.GetSkip(cmd, SKIP),
		reader.GetInitGit(cmd, GIT),
		reader.GetVSCode(cmd, VSCODE),
		reader.IsSetInitGit(cmd, GIT),
		reader.IsSetInitVSCode(cmd, VSCODE),
		models.FromCli,
		generateNewFolder,
	)

	project.Generate(setup)
}
