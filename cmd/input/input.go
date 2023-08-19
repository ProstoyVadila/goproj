package input

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/ProstoyVadila/goproj/internal/models"
)

// Get tries to get information about the project from input.
func Get(configExists bool, conf models.GlobalConfig) (*models.Setup, error) {
	setup := new(models.Setup)
	var useConfig bool

	if err := survey.AskOne(
		packageNameQuestion.Prompt,
		&setup.PackageName,
		survey.WithValidator(survey.Required),
	); err != nil {
		return &models.Setup{}, err
	}

	if configExists {
		configQuestion := getConfigQuestion(conf)
		if err := survey.AskOne(configQuestion.Prompt, &useConfig); err != nil {
			return &models.Setup{}, err
		}
	}

	if !configExists || !useConfig {
		if err := survey.Ask(additionalQsuestions, setup); err != nil {
			return &models.Setup{}, err
		}
	}
	return setup, nil
}
