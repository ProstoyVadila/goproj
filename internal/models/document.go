package models

import (
	"path/filepath"
	"regexp"
	"runtime"
)

type Document struct {
	Name       string
	Path       string
	Filename   string
	Filepath   string
	IsTemplate bool
	// TODO: figure out with the type of Data
	Data any // Document type
}

func NewDocument(name, filename, filepath, templatePath string, isTemplate bool, data any) *Document {
	return &Document{
		Name:       name,
		Path:       templatePath,
		Filename:   filename,
		Filepath:   filepath,
		Data:       data,
		IsTemplate: isTemplate,
	}
}

func (t *Document) FullFilePath() string {
	return filepath.Join(t.Filepath, t.Filename)
}

func (t *Document) FullDocPath() string {
	return filepath.Join(t.Path, t.Name)
}

func GoVersion() string {
	fullVersion := runtime.Version()
	re := regexp.MustCompile(`\d\.\d+`)
	return re.FindString(fullVersion)
}
