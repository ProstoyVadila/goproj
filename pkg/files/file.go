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

func (t *File) get(embedFiles embed.FS) error {
	tmpl, err := template.ParseFS(embedFiles, t.Document.FullDocPath())
	if err != nil {
		return err
	}
	t.Tmpl = tmpl
	return nil
}

func (t *File) create() error {
	file, err := os.Create(t.Document.FullFilePath())
	if err != nil {
		return err
	}
	t.File = file
	return nil
}

func (t *File) write() error {
	err := t.Tmpl.Execute(t.File, t.Document.Data)
	if err != nil {
		return err
	}
	return nil
}

func (t *File) generate(embedFiles embed.FS) error {
	defer t.File.Close()

	err := t.get(embedFiles)
	if err != nil {
		return err
	}
	err = t.create()
	if err != nil {
		return err
	}

	err = t.write()
	if err != nil {
		return err
	}

	return nil
}
