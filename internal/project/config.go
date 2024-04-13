package project

import (
	"github.com/ProstoyVadila/goproj/cmd/input"
	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/internal/models"
)

// StoreConfig
func StoreConfig(argsConf ...*models.GlobalConfig) error {
	conf := new(models.GlobalConfig)

	if len(argsConf) != 0 {
		conf = argsConf[0]
	} else {
		inputSurvey, err := input.Get(false, false)
		if err != nil {
			return err
		}
		conf = inputSurvey.ToGlobalConfig()
	}
	conf.ValidatePrefix()
	conf.Show()
	return config.Store(conf)
}
