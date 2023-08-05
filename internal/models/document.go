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

// FullFilePath returns the full path with the result file name.
func (t *Document) FullFilePath() string {
	return filepath.Join(t.Filepath, t.Filename)
}

// FullDocPath return the full path with the document name.
func (t *Document) FullDocPath() string {
	return filepath.Join(t.Path, t.Name)
}

// GoVersion gets the version of Go from the runtime in the format <major_version>.<minor_version> (1.19 for example).
func GoVersion() string {
	fullVersion := runtime.Version()
	re := regexp.MustCompile(`\d\.\d+`)
	return re.FindString(fullVersion)
}
