package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func getPackageName(args []string) string {
	if len(args) != 0 {
		return args[0]
	}
	return ""
}

func getAuthor(cmd *cobra.Command) string {
	author, err := cmd.Flags().GetString(AUTHOR)
	if err != nil {
		log.Fatal(err)
	}
	return author
}

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

func getFilesToSkip(cmd *cobra.Command) []string {
	filesToSkip, err := cmd.Flags().GetStringSlice(SKIP)
	if err != nil {
		log.Fatal(err)
	}
	if len(filesToSkip) != 0 {
		return filesToSkip
	}
	return make([]string, 0)
}
