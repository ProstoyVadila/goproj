package toml

import (
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/pelletier/go-toml/v2"
)

func Unmarshal(file []byte) (config models.ConfigFromFile, err error) {
	err = toml.Unmarshal(file, &config)
	return
}
