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

var packageCommand = &cobra.Command{
	Use:     "init",
	Short:   "Generate a new Go porject with default files and folders",
	Long:    "Generate a new Go porject with default files (README.md, LICENSE, go.mod, Makefile, Dockerfile, .gitignore, .dockerignore, .env) and folders (cmd/, internal/, pkg/, tests/)",
	Example: "goproj init github.com/Bobert/new_project -a \"Bobert Doe\" -d=\"My new project\" -s=\"Dockerfile,.dockerignore,internal/,pkg/\" --git=false",
	Args:    cobra.ExactArgs(1),
	Run:     packageName,
}

func init() {
	rootCommand.AddCommand(packageCommand)

	// packageCommand.Flags().StringArrayP(AUTHOR)
	packageCommand.Flags().StringP(AUTHOR, "a", "", "an optional flag to set your name")
	packageCommand.Flags().StringSliceP(DESCRIPTION, "d", nil, "an optional flag to set a description of your project")
	packageCommand.Flags().StringSliceP(SKIP, "s", nil, "an optional flag to skip exact files and/or folders (add `/`) from the generation.")
	packageCommand.Flags().BoolP(GIT, "g", true, "an optional flag to define start git initialization or not")
	packageCommand.Flags().BoolP(VSCODE, "c", true, "an optional flag to open the new project in VS Code")
}

// packageName gets poject's setup from CLI and runs generation and initialization of files/git repo.
func packageName(cmd *cobra.Command, args []string) {
	setup := models.NewSetup(
		reader.GetPackageName(args),
		reader.GetAuthor(cmd, AUTHOR),
		reader.GetDescription(cmd, DESCRIPTION),
		reader.GetSkip(cmd, SKIP),
		reader.GetInitGit(cmd, GIT),
		reader.GetVSCode(cmd, VSCODE),
	)

	project.Generate(setup)
}
