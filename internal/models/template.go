package models

import (
	"path/filepath"
	"regexp"
	"runtime"
)

type Document interface {
	ReadmeInfo | LicenseInfo | GoModInfo
}

type TemplateInfo struct {
	Name     string
	Path     string
	Filename string
	Filepath string
	// TODO: figure out with the type of Data
	Data any // Document type
}

func NewTemplateInfo(name, filename, absPath, templatePath string, data any) *TemplateInfo {
	return &TemplateInfo{
		Name:     name,
		Path:     templatePath,
		Filename: filename,
		Filepath: absPath,
		Data:     data,
	}
}

func (t *TemplateInfo) PathWtihFileName() string {
	return filepath.Join(t.Filepath, t.Filename)
}

func (t *TemplateInfo) PathWithTemplateName() string {
	return filepath.Join(t.Path, t.Name)
}

func GoVersion() string {
	fullVersion := runtime.Version()
	re := regexp.MustCompile(`\d\.\d+`)
	return re.FindString(fullVersion)
}
