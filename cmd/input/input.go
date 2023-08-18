package input

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/ProstoyVadila/goproj/internal/models"
)

// GetSetup tries to get information about the project from input.
func GetSetup(configExists bool) (*models.Setup, error) {
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
