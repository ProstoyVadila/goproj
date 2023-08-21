package files

import (
	"embed"

	"github.com/ProstoyVadila/goproj/internal/models"
)

func generateTemplate(info *models.Document, embedFiles embed.FS, errCh chan<- error) {
	template := NewFile(info)
	err := template.Generate(embedFiles)
	if err != nil {
		errCh <- err
	}
	errCh <- nil
}

// Generate creates files from templates and move pre-made files
func Generate(projectInfo *models.ProjectInfo) error {
	errCh := make(chan error, len(projectInfo.Documents))
	defer close(errCh)

	for _, templateInfo := range projectInfo.Documents {
		go generateTemplate(templateInfo, projectInfo.EmbedFiles, errCh)
	}

	for i := 0; i < len(projectInfo.Documents); i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}

	return nil
}
