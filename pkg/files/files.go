package files

import (
	"embed"
	"fmt"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

func generateTemplate(info *models.TemplateInfo, embedFiles embed.FS, errCh chan<- error) {
	template := NewTemplate(info)
	err := template.generate(embedFiles)
	if err != nil {
		errCh <- err
	}
	errCh <- nil
}

func Generate(projectInfo *models.ProjectInfo) error {
	fmt.Println("Generating files for")

	errCh := make(chan error, len(projectInfo.Templates))
	defer close(errCh)

	for _, templateInfo := range projectInfo.Templates {
		go generateTemplate(templateInfo, projectInfo.EmbedFiles, errCh)
	}

	for i := 0; i < len(projectInfo.Templates); i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}

	return nil
}
