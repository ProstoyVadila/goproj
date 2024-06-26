package input

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/output"
)

// TODO; refactor this crap
// Get tries to get information through survey in input.
func Get(configExists, isInit bool, conf ...*models.GlobalConfig) (*models.Survey, error) {
	output.Info("\nLet's start!")
	surv := new(models.Survey)
	var useConfig bool

	if isInit {
		packageNameQuestion := getPackageNameQuestion(configExists, conf...)
		err := survey.AskOne(
			packageNameQuestion.Prompt,
			&surv.PackageName,
			survey.WithValidator(survey.Required),
		)
		additionalQsuestions = append(additionalQsuestions, descriptionQuestion)
		if err != nil {
			return surv, err
		}
	}

	if configExists && len(conf) != 0 {
		configQuestion := getConfigQuestion(conf[0])
		if err := survey.AskOne(configQuestion.Prompt, &useConfig); err != nil {
			return surv, err
		}
	}

	if !configExists || !useConfig {
		if !isInit {
			additionalQsuestions = append(additionalQsuestions, prefixQuestion)
		}
		if err := survey.Ask(additionalQsuestions, surv); err != nil {
			return surv, err
		}
	}
	return surv, nil
}
