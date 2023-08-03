package files

import (
	"fmt"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

func generateTemplate(info *models.TemplateInfo, errCh chan<- error) {
	template := NewTemplate(info)
	err := template.generate()
	if err != nil {
		errCh <- err
	}
	errCh <- nil
}

func Generate(projectInfo *models.ProjectInfo) error {
	fmt.Printf("Generating files for %s\n", projectInfo.AuthorName)

	errCh := make(chan error)
	defer close(errCh)

	for _, templateInfo := range projectInfo.Templates {
		go generateTemplate(templateInfo, errCh)
	}

	for i := 0; i < len(projectInfo.Templates); i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}
	return nil
}
