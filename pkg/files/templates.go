package files

import (
	"os"
	"text/template"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

type Templapte struct {
	File *os.File
	Tmpl *template.Template
	*models.TemplateInfo
}

func NewTemplate(templateInfo *models.TemplateInfo) *Templapte {
	return &Templapte{
		TemplateInfo: templateInfo,
	}
}

func (t *Templapte) getTemplateFile() error {
	tmpl, err := template.ParseFiles(t.TemplateInfo.PathWithTemplateName())
	if err != nil {
		return err
	}
	t.Tmpl = tmpl
	return nil
}

func (t *Templapte) createFile() error {
	file, err := os.Create(t.TemplateInfo.PathWtihFileName())
	if err != nil {
		return err
	}
	t.File = file
	return nil
}

func (t *Templapte) writeFile() error {
	err := t.Tmpl.Execute(t.File, t.TemplateInfo.Data)
	if err != nil {
		return err
	}
	return nil
}

func (t *Templapte) generate() error {
	defer t.File.Close()

	err := t.getTemplateFile()
	if err != nil {
		return err
	}
	err = t.createFile()
	if err != nil {
		return err
	}
	err = t.writeFile()
	if err != nil {
		return err
	}
	return nil
}