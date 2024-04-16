package models

import (
	"path/filepath"
	"strings"

	"github.com/ProstoyVadila/goproj/pkg/output"
	"github.com/elliotchance/orderedmap/v2"
)

const (
	FromCli = iota + 1
	FromConfig
	FromSurvey
)

type Setup struct {
	Skip            []string
	PackageName     string
	MainFolder      string
	Author          string
	Description     string
	InitGit         bool
	InitVSCode      bool
	IsSetInitGit    bool
	IsSetInitVSCode bool
	From            int
}

// NewSetup constructs Setup by fields.
func NewSetup(
	packageName,
	author,
	description string,
	skip []string,
	initGit,
	initVSCode,
	isSetInitGit,
	isSetInitVSCode bool,
	from int,
) *Setup {

	mainFolder := validateMainFolder(packageName)
	return &Setup{
		PackageName:     packageName,
		MainFolder:      mainFolder,
		Author:          author,
		Description:     description,
		Skip:            skip,
		InitGit:         initGit,
		InitVSCode:      initVSCode,
		IsSetInitGit:    isSetInitGit,
		IsSetInitVSCode: isSetInitVSCode,
		From:            from,
	}
}

// NewSetupFromConfig costructs Setup from ConfigFromFile.
func NewSetupFromConfig(conf *GlobalConfig) *Setup {
	return &Setup{
		Author:          conf.Author,
		Skip:            conf.Skip,
		InitGit:         conf.InitGit,
		InitVSCode:      conf.InitVSCode,
		IsSetInitGit:    true,
		IsSetInitVSCode: true,
		From:            FromConfig,
	}
}

// Update updates Setup fields by another Setup
func (s *Setup) Update(from *Setup) {
	s.PackageName = from.PackageName
	s.MainFolder = from.MainFolder

	if from.Author != "" {
		s.Author = from.Author
	}
	if from.Description != "" {
		s.Description = from.Description
	}
	if len(from.Skip) != 0 {
		s.Skip = from.Skip
	}
	s.updateInitGit(from)
	s.updateInitVSCode(from)
}

// updateInitGit sets InitGit if it was set in fromSetup
func (s *Setup) updateInitGit(from *Setup) {
	if from.IsSetInitGit {
		s.InitGit = from.InitGit
	}
}

// updateInitVSCode sets InitVSCode if it was set in fromSetup
func (s *Setup) updateInitVSCode(from *Setup) {
	if from.IsSetInitVSCode {
		s.InitVSCode = from.InitVSCode
	}
}

// ValidatePackageName adds prefix from global config
func (s *Setup) ValidatePackageName(prefix string) {
	if strings.Contains(s.PackageName, prefix) || strings.Contains(s.PackageName, "github") {
		return
	}
	s.PackageName = filepath.Join(prefix, s.PackageName)
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

// getShow creates ordered map of Setup fields and msg for output.
func (s *Setup) getShow() (*orderedmap.OrderedMap[string, any], string) {
	msg := "This your project setup:"
	omap := orderedmap.NewOrderedMap[string, any]()

	omap.Set("Project (package) name: %s", s.PackageName)
	omap.Set("Author: %s", s.Author)
	omap.Set("Description: %s", s.Description)
	omap.Set("Files to skip: %v", s.FilesToSkip())
	omap.Set("Folders to skip: %v", s.FoldersToSkip())
	omap.Set("Init Git Repo: %v", s.InitGit)
	omap.Set("Open in VS Code: %v", s.InitVSCode)

	return omap, msg
}

// Show writes Setup info to standart output.
func (s *Setup) Show() {
	output.Show(s.getShow())
}

// ShowString returns output string for Setup.
func (s *Setup) ShowString() string {
	return output.ShowString(s.getShow())
}

func validateMainFolder(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimRight(name, "/")
	name_items := strings.Split(name, "/")
	if len(name_items) == 1 {
		return name
	}
	return name_items[len(name_items)-1]
}
