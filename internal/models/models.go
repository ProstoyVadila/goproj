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

type ProjectInfo struct {
	Templates   []*TemplateInfo
	EmbedFiles  embed.FS
	AuthorName  string
	PackageName string
	Path        string
}

func (p *ProjectInfo) TemplatePath() string {
	// return filepath.Join(p.Path, TEMPLATE_PATH)
	return TEMPLATE_PATH
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

	projectInfo.AddTemplate(license, readme, gomod)

	return projectInfo
}

func (p *ProjectInfo) AddTemplate(templates ...*TemplateInfo) {
	p.Templates = append(p.Templates, templates...)
}
