package models

import (
	"path/filepath"
)

type Document struct {
	DataToAdd    []string
	Constructor  any
	Data         any
	Name         string
	TemplatePath string
	Filename     string
	Filepath     string
	IsTemplate   bool
}

func NewDocument(name, filename, templatePath string, constructor any, isTemplate bool, dataToAdd []string) *Document {
	return &Document{
		Name:         name,
		TemplatePath: templatePath,
		Filename:     filename,
		Constructor:  constructor,
		IsTemplate:   isTemplate,
		DataToAdd:    dataToAdd,
	}
}

// FullFilePath returns the full path with the result file name.
func (t *Document) FullFilePath() string {
	return filepath.Join(t.Filepath, t.Filename)
}

// FullDocPath return the full path with the document name.
func (t *Document) FullDocPath() string {
	return filepath.Join(t.TemplatePath, t.Name)
}
