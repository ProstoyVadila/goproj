package folders

import (
	"fmt"
	"os"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

func Create(folders []*models.Folder) error {
	fmt.Println("Generating folders")

	for _, folder := range folders {
		err := os.Mkdir(folder.Name, folder.Permissions)
		if err != nil {
			return err
		}
	}
	return nil
}
