package git

import (
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/cmd"
)

const (
	gitCommand = "git"
	arg0       = "init"
)

// InitGitRepo initialize git repository
func InitGitRepo(projectInfo *models.ProjectInfo) error {
	command := cmd.New(gitCommand, arg0)
	if err := command.Execute(); err != nil {
		return err
	}
	return nil
}
