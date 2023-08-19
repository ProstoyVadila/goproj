package models

import "fmt"

type Setup struct {
	Skip        []string
	PackageName string
	Author      string
	Description string
	InitGit     bool
	InitVSCode  bool
}

// NewSetup constructs Setup by fields.
func NewSetup(packageName, author, description string, skip []string, skipGit, initVSCode bool) *Setup {
	return &Setup{
		PackageName: packageName,
		Author:      author,
		Description: description,
		Skip:        skip,
		InitGit:     skipGit,
		InitVSCode:  initVSCode,
	}
}

// NewSetupFromConfig costructs Setup from ConfigFromFile.
func NewSetupFromConfig(conf GlobalConfig) *Setup {
	return &Setup{
		Author:      conf.Author,
		Description: conf.Description,
		Skip:        conf.Skip,
		InitGit:     conf.InitGit,
		InitVSCode:  conf.InitVSCode,
	}
}

// Update updates Setup fields by another Setup
func (s *Setup) Update(from *Setup) {
	s.PackageName = from.PackageName
	if from.Author != "" {
		s.Author = from.Author
	}
	if from.Description != "" {
		s.Description = from.Description
	}
	if len(from.Skip) != 0 {
		s.Skip = from.Skip
	}
	if from.InitGit || s.InitGit != from.InitGit {
		s.InitGit = from.InitGit
	}
	if from.InitVSCode || s.InitVSCode != from.InitVSCode {
		s.InitVSCode = from.InitVSCode
	}
}

// FilesToSkip gets files from skip objects.
func (s *Setup) FilesToSkip() []string {
	var files []string
	for _, object := range s.Skip {
		if object[len(object)-1] != '/' {
			files = append(files, object)
		}
	}
	return files
}

// FoldersToSkip gets folders from skip objects.
func (s *Setup) FoldersToSkip() []string {
	var folders []string
	for _, object := range s.Skip {
		last := len(object) - 1
		if object[last] == '/' {
			folders = append(folders, object[:last])
		}
	}
	return folders
}

// showSetup writes setup info from CLI to standart output.
func (s *Setup) Show() {
	fmt.Printf("\nProject (package) name: %s\n", s.PackageName)
	fmt.Printf("Author: %s\n", s.Author)
	fmt.Printf("Description: %s\n", s.Description)
	fmt.Printf("Files to skip: %v\n", s.FilesToSkip())
	fmt.Printf("Folders to skip: %v\n", s.FoldersToSkip())
	fmt.Printf("Init Git Repo: %v\n", s.InitGit)
	fmt.Printf("Open in VS Code: %v\n\n", s.InitVSCode)
}
