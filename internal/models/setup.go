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

func NewSetupFromConfig(conf ConfigFromFile, filesToSkip, foldersToSkip []string) *Setup {
	return &Setup{
		Author:        conf.Author,
		Description:   conf.Description,
		FilesToSkip:   filesToSkip,
		FoldersToSkip: foldersToSkip,
		InitGit:       conf.InitGit,
		InitVSCode:    conf.InitVSCode,
	}
}

func UpdateSetup(setup1, setup2 *Setup) {
	setup1.PackageName = setup2.PackageName
	if setup2.Author != "" {
		setup1.Author = setup2.Author
	}
	if setup2.Description != "" {
		setup1.Description = setup2.Description
	}
	if len(setup2.FilesToSkip) != 0 {
		setup1.FilesToSkip = setup2.FilesToSkip
	}
	if len(setup2.FoldersToSkip) != 0 {
		setup1.FoldersToSkip = setup2.FilesToSkip
	}
	if !setup2.InitGit {
		setup1.InitGit = setup2.InitGit
	}
	if !setup2.InitVSCode {
		setup1.InitVSCode = setup2.InitVSCode
	}
}
