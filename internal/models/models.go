package models

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"time"
)

// const TEMPLATE_PATH = "/Users/vadim/Documents/scripts/groproj_script/templates"
const (
	TEMPLATE_PATH        = "./templates"
	LOCALRUN_RESULT_PATH = "./tests/tempFiles"

	LICENSE_TEMPLATE = "LICENSE.tmpl"
	LICENSE_FILE     = "LICENSE"
	README_TEMPLATE  = "README.tmpl"
	README_FILE      = "README.md"
)

type TemplateInfo struct {
	Name     string
	Path     string
	Filename string
	Filepath string
	Data     any
}

type LicenseInfo struct {
	Version     string
	AuthorName  string
	LicenseFile string
	Year        int
}

type ReadmeInfo struct {
	AuthorName  string
	Description string
}

type ProjectInfo struct {
	Templates   []*TemplateInfo
	EmbedFiles  embed.FS
	AuthorName  string
	PackageName string
	Path        string
}

func (t *TemplateInfo) PathWtihFileName() string {
	return filepath.Join(t.Filepath, t.Filename)
}

func (t *TemplateInfo) PathWithTemplateName() string {
	return filepath.Join(t.Path, t.Name)
}

func (p *ProjectInfo) TemplatePath() string {
	// return filepath.Join(p.Path, TEMPLATE_PATH)
	return TEMPLATE_PATH
}

// TODO: get path from os
func NewTemplateInfo(name, filename, absPath, templatePath string, data any) *TemplateInfo {
	return &TemplateInfo{
		Name:     name,
		Path:     templatePath,
		Filename: filename,
		Filepath: absPath,
		Data:     data,
	}
}

func NewLicenceInfo(authorName string) *LicenseInfo {
	return &LicenseInfo{AuthorName: authorName, Year: time.Now().Year()}
}

func NewReadmeInfo(authorName, description string) *ReadmeInfo {
	return &ReadmeInfo{AuthorName: authorName, Description: description}
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

	projectInfo.Templates = append(projectInfo.Templates, license, readme)

	return projectInfo
}
