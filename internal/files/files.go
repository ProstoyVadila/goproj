package files

import (
	"fmt"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

func generateTemplate(info *models.TemplateInfo) error {
	licenseTemplate := NewTemplate(info)
	err := licenseTemplate.generate()
	if err != nil {
		return err
	}
	return nil
}

func Generate(projectInfo *models.ProjectInfo) error {
	fmt.Printf("Generating files for %s\n", projectInfo.AuthorName)

	for _, templateInfo := range projectInfo.Templates {
		if err := generateTemplate(templateInfo); err != nil {
			return err
		}
	}

	return nil
}

func generateTemplate2(info *models.TemplateInfo, errCh chan<- error) {
	licenseTemplate := NewTemplate(info)
	err := licenseTemplate.generate()
	if err != nil {
		errCh <- err
	}
	errCh <- nil
}

func Generate2(projectInfo *models.ProjectInfo) error {
	fmt.Printf("Generating files for %s\n", projectInfo.AuthorName)

	errCh := make(chan error)
	defer close(errCh)

	for _, templateInfo := range projectInfo.Templates {
		go generateTemplate2(templateInfo, errCh)
	}

	for i := 0; i < len(projectInfo.Templates); i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}
	return nil
}
