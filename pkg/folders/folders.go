package folders

import (
	"os"

	"github.com/ProstoyVadila/goproj/internal/models"
)

// Create creates folders
func Create(folders []*models.Folder) error {
	for _, folder := range folders {
		err := os.Mkdir(folder.Name, folder.Permissions)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateOne create folder
func CreateOne(foders *models.Folder) error {
	return os.Mkdir(foders.Name, foders.Permissions)
}
