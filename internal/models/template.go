package models

import "path/filepath"

type TemplateInfo struct {
	Name     string
	Path     string
	Filename string
	Filepath string
	Data     any
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
