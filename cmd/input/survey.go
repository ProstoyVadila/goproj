package input

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/internal/models"
)

// packageNameQuestion is a required question about the new generated project name
var packageNameQuestion = &survey.Question{
	Name: "PackageName",
	Prompt: &survey.Input{
		Message: "Project (package) name:",
		Help:    "For example: github.com/me/new_project",
	},
	Validate: survey.Required,
}

// descriptionQuestion
var descriptionQuestion = &survey.Question{
	Name: "Description",
	Prompt: &survey.Input{
		Message: "Description",
		Help:    "You can set a small description in README.md for your new projct.",
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
