package git

import (
	"github.com/ProstoyVadila/goprojtemplate/internal/models"
	"github.com/ProstoyVadila/goprojtemplate/pkg/cmd"
)

const gitCommand = "git"
const arg0 = "init"

func InitGitRepo(projectInfo *models.ProjectInfo) error {
	command := cmd.New(gitCommand, arg0)
	if err := command.Execute(); err != nil {
		return err
	}
	return nil
}
