package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestDocument() *Document {
	name := "Dockerfile.tmpl"
	templatePath := "templates"
	filename := "Dockerfile"
	constructor := struct{}{}
	isTemplate := true
	filepath := "example"
	dataToAdd := []string{"smth"}

	doc := &Document{
		Name:         name,
		TemplatePath: templatePath,
		Filename:     filename,
		Constructor:  constructor,
		IsTemplate:   isTemplate,
		DataToAdd:    dataToAdd,
		Filepath:     filepath,
	}
	return doc
}

func TestNewDocument(t *testing.T) {
	doc1 := getTestDocument()
	doc2 := NewDocument(
		doc1.Name,
		doc1.Filename,
		doc1.TemplatePath,
		doc1.Constructor,
		doc1.IsTemplate,
		doc1.DataToAdd,
	)
	doc2.Filepath = doc1.Filepath

	assert.Equal(t, doc1, doc2)
}

func TestDocumentFullFilePath(t *testing.T) {
	doc := getTestDocument()

	slash := "/"
	filepath1 := doc.Filepath + slash + doc.Filename
	filepath2 := doc.Filepath + slash + slash + doc.Filename

	assert.Equal(t, filepath1, doc.FullFilePath())
	assert.NotEqual(t, filepath2, doc.FullFilePath())
}

func TestDocumentFullDocPath(t *testing.T) {
	doc := getTestDocument()

	slash := "/"
	docfilepath1 := doc.TemplatePath + slash + doc.Name
	docfilepath2 := doc.TemplatePath + slash + slash + doc.Name

	assert.Equal(t, docfilepath1, doc.FullDocPath())
	assert.NotEqual(t, docfilepath2, doc.FullDocPath())
}
