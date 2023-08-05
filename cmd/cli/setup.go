package cli

import (
	"fmt"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/project"
	"github.com/spf13/cobra"
)

var packageCommand = &cobra.Command{
	Use:   "init",
	Short: "Your new package name",
	Long:  "Your new package name, like 'github.com/bla/blabla'",
	Run:   packageName,
}

func init() {
	rootCommand.AddCommand(packageCommand)

	packageCommand.PersistentFlags().StringP(AUTHOR, "a", "", "An optional flag to set your name")
	packageCommand.PersistentFlags().StringSliceP(DESCRIPTION, "d", nil, "An optional flag to set description of your project")
	packageCommand.PersistentFlags().StringSliceP(SKIP, "s", nil, "A flag to skip exact files")
}

// packageName gets poject's setup from CLI and runs generation and initialization of files/git repo.
func packageName(cmd *cobra.Command, args []string) {
	fmt.Println("I get package name")

	setup := models.NewSetup(
		getPackageName(args),
		getAuthor(cmd),
		getDescription(cmd),
		getSkip(cmd),
	)

	showSetup(setup)
	project.Generate(setup)
}

// showSetup writes setup info from CLI to standart output.
func showSetup(setup *models.Setup) {
	fmt.Printf("\nProject (package) name: %s\n", setup.PackageName)
	fmt.Printf("Author: %s\n", setup.Author)
	fmt.Printf("Description: %s\n", setup.Description)
	fmt.Printf("Files to skip: %v\n", setup.FilesToSkip)
}
