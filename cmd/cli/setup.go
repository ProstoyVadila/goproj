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

func packageName(cmd *cobra.Command, args []string) {
	fmt.Println("I get package name")

	dataCli := models.NewInputData(
		getPackageName(args),
		getAuthor(cmd),
		getDescription(cmd),
		getSkip(cmd),
	)

	fmt.Println("DataCLI")
	fmt.Printf("package name: %s\n", dataCli.PackageName)
	fmt.Printf("author: %s\n", dataCli.Author)
	fmt.Printf("descr: %s\n", dataCli.Description)
	fmt.Printf("skip: %v\n", dataCli.FilesToSkip)

	project.Generate(dataCli)
}
