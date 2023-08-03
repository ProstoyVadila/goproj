package git

import (
	"fmt"
	"os/exec"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

const command = "git"
const arg0 = "init"

func InitGitRepo(projectInfo *models.ProjectInfo) error {
	fmt.Println("Initing git repo")
	cmd := exec.Command(command, arg0)
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(stdout)
	return nil
}
