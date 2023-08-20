package yaml

import (
	"github.com/ProstoyVadila/goproj/internal/models"
	"gopkg.in/yaml.v3"
)

func Unmarshal(file []byte) (config *models.GlobalConfig, err error) {
	err = yaml.Unmarshal(file, &config)
	return
}
