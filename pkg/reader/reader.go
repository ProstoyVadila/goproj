package reader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/reader/json"
	"github.com/ProstoyVadila/goproj/pkg/reader/toml"
	"github.com/ProstoyVadila/goproj/pkg/reader/yaml"
)

type Reader interface {
	Unmarshal(file []byte) (*models.ConfigFromFile, error)
}

// Unmarshal reads file provided in CLI to models.ConfigFromFile
func ConfigFromFile(filename string) (config models.ConfigFromFile, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	switch filepath.Ext(filename) {
	case ".json":
		return json.Unmarshal(file)
	case ".toml":
		return toml.Unmarshal(file)
	case ".yaml":
		return yaml.Unmarshal(file)
	default:
		return config, fmt.Errorf("unknown file extension: %s", filename)
	}
}
