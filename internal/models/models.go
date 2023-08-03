package models

import (
	"path/filepath"
	"time"
)

const TEMPLATE_PATH = "./templates"
const RESULT_PATH = "./tests/tempFiles"

const LICENSE_TEMPLATE = "LICENSE.tmpl"
const LICENSE_FILE = "LICENSE"
const README_TEMPLATE = "README.tmpl"
const README_FILE = "README.md"

type TemplateInfo struct {
	Name     string
	Path     string
	Filename string
	Filepath string
	Data     any
}

func (t *TemplateInfo) PathWtihFileName() string {
	return filepath.Join(t.Filepath, t.Filename)
}

func (t *TemplateInfo) PathWithTemplateName() string {
	return filepath.Join(t.Path, t.Name)
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
	Templates    []*TemplateInfo
	AuthorName   string
	PackageName  string
	TemplatePath string
}

// TODO: get path from os
func NewTemplateInfo(name, filename string, data any) *TemplateInfo {
	return &TemplateInfo{
		Name:     name,
		Path:     TEMPLATE_PATH,
		Filename: filename,
		Filepath: RESULT_PATH,
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
	projectInfo := &ProjectInfo{
		AuthorName:   authorName,
		PackageName:  packageName,
		TemplatePath: TEMPLATE_PATH,
	}

	license := NewTemplateInfo(
		LICENSE_TEMPLATE,
		LICENSE_FILE,
		NewLicenceInfo(authorName),
	)
	readme := NewTemplateInfo(
		README_TEMPLATE,
		README_FILE,
		NewReadmeInfo(authorName, description),
	)

	projectInfo.Templates = append(projectInfo.Templates, license, readme)

	return projectInfo
}
