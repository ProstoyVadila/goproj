package models

type Setup struct {
	FilesToSkip   []string
	FoldersToSkip []string
	PackageName   string
	Author        string
	Description   string
	InitGit       bool
	InitVSCode    bool
}

func NewSetup(packageName, author, description string, filesToSkip, foldersToSkip []string, skipGit, initVSCode bool) *Setup {
	return &Setup{
		PackageName:   packageName,
		Author:        author,
		Description:   description,
		FilesToSkip:   filesToSkip,
		FoldersToSkip: foldersToSkip,
		InitGit:       skipGit,
		InitVSCode:    initVSCode,
	}
}
