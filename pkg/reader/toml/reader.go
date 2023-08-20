package toml

import (
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/pelletier/go-toml/v2"
)

func Unmarshal(file []byte) (config models.GlobalConfig, err error) {
	err = toml.Unmarshal(file, &config)
	return
}

func Marshal(config models.GlobalConfig) (data []byte, err error) {
	return toml.Marshal(config)
}
