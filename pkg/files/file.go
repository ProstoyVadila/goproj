package files

import (
	"embed"
	"os"
	"text/template"

	"github.com/ProstoyVadila/goproj/internal/models"
)

type File struct {
	File *os.File
	Tmpl *template.Template
	*models.Document
}

func NewFile(doc *models.Document) *File {
	return &File{
		Document: doc,
	}
}

// Get parses the file (or the template file) from the embed files
func (t *File) Get(embedFiles embed.FS) error {
	tmpl, err := template.ParseFS(embedFiles, t.Document.FullDocPath())
	if err != nil {
		return err
	}
	t.Tmpl = tmpl
	return nil
}

// Create creates file in the path
func (t *File) Create() error {
	file, err := os.Create(t.Document.FullFilePath())
	if err != nil {
		return err
	}
	t.File = file
	return nil
}

// Write writes tempate data to file
func (t *File) Write() error {
	return t.Tmpl.Execute(t.File, t.Document.Data)
}

// Generate creates new files from the embed files with data if it's a template
func (t *File) Generate(embedFiles embed.FS) (err error) {
	defer t.File.Close()

	err = t.Get(embedFiles)
	if err != nil {
		return
	}
	err = t.Create()
	if err != nil {
		return
	}

	return t.Write()
}
