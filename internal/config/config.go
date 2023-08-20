package config

import (
	"log"
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

func Store(config *models.GlobalConfig) (err error) {
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

func Get() (config *models.GlobalConfig, ok bool) {
	if !configExists(ConfigFilepath) {
		return
	}
	conf, err := reader.GetGlobalConfig(ConfigFilepath)
	if err != nil {
		log.Fatal(err)
	}
	return conf, true
}
