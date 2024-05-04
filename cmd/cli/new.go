package cli

import (
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/project"
	"github.com/ProstoyVadila/goproj/pkg/reader"
	"github.com/spf13/cobra"
)

const (
	SKIP        = "skip"
	AUTHOR      = "author"
	DESCRIPTION = "description"
	GIT         = "git"
	VSCODE      = "vscode"
)

var newPackageCommand = &cobra.Command{
	Use:     "new",
	Short:   "Generate a new Go porject in the new folder with default files and folders",
	Long:    "Generate a new Go porject in the new folder with default files (README.md, LICENSE, go.mod, Makefile, Dockerfile, .gitignore, .dockerignore, .env) and folders (cmd/, internal/, pkg/, tests/)",
	Example: "goproj new github.com/Bobert/new_project -a \"Bobert Doe\" -d=\"My new project\" -s=\"Dockerfile,.dockerignore,internal/,pkg/\" --git=false",
	Args:    cobra.ExactArgs(1),
	Run:     newPackageName,
}

func init() {
	rootCommand.AddCommand(newPackageCommand)

	newPackageCommand.Flags().StringP(AUTHOR, "a", "", "an optional flag to set your name")
	newPackageCommand.Flags().StringSliceP(DESCRIPTION, "d", nil, "an optional flag to set a description of your project")
	newPackageCommand.Flags().StringSliceP(SKIP, "s", nil, "an optional flag to skip exact files and/or folders (add `/` after folder's name) from the generation.")
	newPackageCommand.Flags().BoolP(GIT, "g", false, "an optional flag to define start git initialization or not (Default false)")
	newPackageCommand.Flags().BoolP(VSCODE, "c", false, "an optional flag to open the new project in VS Code (Default false)")
}

// newPackageName gets poject's setup from CLI and runs generation and initialization of files/git repo.
func newPackageName(cmd *cobra.Command, args []string) {
	const generateNewFolder = true
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
