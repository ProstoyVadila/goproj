package project

import (
	"log"

	"github.com/ProstoyVadila/goproj/cmd/input"
	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/internal/models"
)

// enrichSetup updates Setup with GlobalConfig and/or CLI args, Input
func enrichSetup(ArgsSetup ...*models.Setup) *models.Setup {
	setup := new(models.Setup)

	// Get setup from the global config
	conf, confExists := config.Get()
	if confExists {
		setup = models.NewSetupFromConfig(conf)
	}

	// Update with CLI args or Input
	if len(ArgsSetup) == 1 {
		setup.Update(ArgsSetup[0])
	} else {
		inputSurvey, err := input.Get(confExists, true, conf)
		if err != nil {
			log.Fatal(err)
		}
		setup.Update(inputSurvey.ToSetup())
	}

	return setup
}
