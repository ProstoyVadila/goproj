package models

import (
	"embed"
	"log"
	"os"
)

const (
	TEMPLATE_PATH        = "./templates"
	LOCALRUN_RESULT_PATH = "./tests/tempFiles"
)

// TODO get values from input/setup
var LIST_OF_FOLDERS = [3]string{"test_cmd", "test_pkg", "test_internal"}

type ProjectInfo struct {
	Templates   []*TemplateInfo
	EmbedFiles  embed.FS
	AuthorName  string
	PackageName string
	Path        string
	InitGit     bool
	Folders     []*Folder
}

func (p *ProjectInfo) TemplatePath() string {
	// return filepath.Join(p.Path, TEMPLATE_PATH)
	return TEMPLATE_PATH
}

func (p *ProjectInfo) AddTemplates(templates ...*TemplateInfo) {
	p.Templates = append(p.Templates, templates...)
}

func (p *ProjectInfo) AddFolders(folders ...*Folder) {
	p.Folders = append(p.Folders, folders...)
}

func NewProjectInfo(authorName, packageName, description string) *ProjectInfo {
	absPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectInfo := &ProjectInfo{
		AuthorName:  authorName,
		PackageName: packageName,
		Path:        absPath,
		// TODO get InitGit value from input questions
		InitGit: true,
	}

	license := NewTemplateInfo(
		LICENSE_TEMPLATE,
		LICENSE_FILE,
		absPath,
		projectInfo.TemplatePath(),
		NewLicenceInfo(authorName),
	)
	readme := NewTemplateInfo(
		README_TEMPLATE,
		README_FILE,
		absPath,
		projectInfo.TemplatePath(),
		NewReadmeInfo(authorName, description),
	)
	gomod := NewTemplateInfo(
		GOMOD_TEMPLATE,
		GOMOD_FILE,
		absPath,
		projectInfo.TemplatePath(),
		NewGoModInfo(packageName),
	)
	projectInfo.AddTemplates(license, readme, gomod)

	for _, folderName := range LIST_OF_FOLDERS {
		projectInfo.AddFolders(NewFolder(folderName))
	}
	return projectInfo
}
