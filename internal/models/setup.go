package models

type Setup struct {
	Skip        []string
	PackageName string
	Author      string
	Description string
	InitGit     bool
	InitVSCode  bool
}

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

func NewSetupFromConfig(conf ConfigFromFile) *Setup {
	return &Setup{
		Author:      conf.Author,
		Description: conf.Description,
		Skip:        conf.Skip,
		InitGit:     conf.InitGit,
		InitVSCode:  conf.InitVSCode,
	}
}

func (s *Setup) Update(anotherSetup *Setup) {
	s.PackageName = anotherSetup.PackageName
	if anotherSetup.Author != "" {
		s.Author = anotherSetup.Author
	}
	if anotherSetup.Description != "" {
		s.Description = anotherSetup.Description
	}
	if len(anotherSetup.Skip) != 0 {
		s.Skip = anotherSetup.Skip
	}
	if anotherSetup.InitGit {
		s.InitGit = anotherSetup.InitGit
	}
	if anotherSetup.InitVSCode {
		s.InitVSCode = anotherSetup.InitVSCode
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
