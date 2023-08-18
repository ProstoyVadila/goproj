package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/reader"
	"github.com/ProstoyVadila/goproj/pkg/reader/toml"
)

const ConfigName = ".goproj.conf.toml"

var (
	UserHomeDir, _ = os.UserHomeDir()
	ConfigFilepath = filepath.Join(UserHomeDir, ConfigName)
)

func configExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func Store(config models.ConfigFromFile) (err error) {
	file, err := os.Create(ConfigFilepath)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := toml.Marshal(config)
	if err != nil {
		return
	}

	_, err = file.Write(data)
	return
}

func Get() (config models.ConfigFromFile, err error) {
	if !configExists(ConfigFilepath) {
		return config, errors.New("cannot find configuration file")
	}
	return reader.ConfigFromFile(ConfigFilepath)
}
