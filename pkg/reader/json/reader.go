package json

import (
	"encoding/json"

	"github.com/ProstoyVadila/goproj/internal/models"
)

func Unmarshal(file []byte) (config models.GlobalConfig, err error) {
	err = json.Unmarshal(file, &config)
	return
}
