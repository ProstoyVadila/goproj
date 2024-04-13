package input

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/internal/models"
)

// descriptionQuestion
var descriptionQuestion = &survey.Question{
	Name: "Description",
	Prompt: &survey.Input{
		Message: "Description:",
		Help:    "You can set a small description for README.md in your new projct.",
	},
}

var prefixHelp = `
This prefix will be added to package name in your future projects. 
Example of prefix: github.com/<your_github_name>
`
var prefixQuestion = &survey.Question{
	Name: "Prefix",
	Prompt: &survey.Input{
		Message: "Projects' global prefix: github.com/",
		Help:    prefixHelp,
	},
}

// additionalQuestions variable is additional input questions
var additionalQsuestions = []*survey.Question{
	{
		Name: "Author",
		Prompt: &survey.Input{
			Message: "Tell me your name:",
			Help:    "It will be set in the LICENSE.",
		},
		Validate: survey.MaxLength(255),
	},

	{
		Name: "Skip",
		Prompt: &survey.MultiSelect{
			Message: "Skip files/folders from generation:",
			Options: []string{
				"cmd/",
				"pkg/",
				"internal/",
				"Makefile",
				"Dockerfile",
				".dockerignore",
				".gitignore",
				"README.md",
				"LICENSE",
				".env",
				"main.go",
				"go.mod",
			},
			PageSize: 12,
			Help:     "You can skip some default files and/or folders from generation fro this project.",
		},
	},
	{
		Name: "InitGit",
		Prompt: &survey.Confirm{
			Message: "Do you want generator to init git repo in this folder?",
			Default: true,
		},
	},
	{
		Name: "InitVSCode",
		Prompt: &survey.Confirm{
			Message: "Do you want generator to open the new project in VS Code?",
			Default: true,
		},
	},
}

func getPackageNameQuestion(confExists bool, confs ...*models.GlobalConfig) *survey.Question {
	message := "Project (package) name:"
	help := "For example: github/blabla/new_project"
	if confExists && len(confs) != 0 {
		message += " " + confs[0].Prefix
		help = fmt.Sprintf("Your prefix %s will be added your project name", confs[0].Prefix)
	}
	return &survey.Question{
		Name: "PackageName",
		Prompt: &survey.Input{
			Message: message,
			Help:    help,
		},
	}
}

// getConfigQuestion creates a question (type Confirm) about using GlobalConfig or not and provides it in the Question's Help.
func getConfigQuestion(conf *models.GlobalConfig) *survey.Question {
	help := fmt.Sprintf("Your global config is located in the \"~/%s\" file. More info via \"goproj config --help\" command.\n", config.ConfigName)
	return &survey.Question{
		Name: "config",
		Prompt: &survey.Confirm{
			Message: "Do you want to use a global config?",
			Help:    conf.ShowString() + help,
			Default: true,
		},
	}
}
