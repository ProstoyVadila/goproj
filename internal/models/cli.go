package models

type Setup struct {
	FilesToSkip   []string
	FoldersToSkip []string
	PackageName   string
	Author        string
	Description   string
}

func NewSetup(packageName, author, description string, filesToSkip, foldersToSkip []string) *Setup {
	return &Setup{
		PackageName:   packageName,
		Author:        author,
		Description:   description,
		FilesToSkip:   filesToSkip,
		FoldersToSkip: foldersToSkip,
	}
}
