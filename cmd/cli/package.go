package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var packageCommand = &cobra.Command{
	Use:   "create",
	Short: "Your new package name",
	Long:  "Your new package name, like 'github.com/bla/blabla'",
	Run:   packageName,
}

func init() {
	rootCommand.AddCommand(packageCommand)

	packageCommand.PersistentFlags().StringSliceP("skip", "s", nil, "A flag to skip exact files")
	packageCommand.PersistentFlags().StringP("author", "a", "", "Optional flag to set your name")
}

func packageName(cmd *cobra.Command, args []string) {
	fmt.Println("I get package name")
	fmt.Println(args)
	skip, _ := cmd.Flags().GetStringSlice("skip")
	if len(skip) != 0 {
		for _, fileToSkip := range skip {
			fmt.Println(fileToSkip)
		}
	}
	author, _ := cmd.Flags().GetString("author")
	if author != "" {
		fmt.Println(author)
	}

}
