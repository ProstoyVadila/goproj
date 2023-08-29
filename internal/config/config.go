package config

import (
	"os"
	"path/filepath"

	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/folders"
	"github.com/ProstoyVadila/goproj/pkg/output"
	"github.com/ProstoyVadila/goproj/pkg/reader"
	"github.com/ProstoyVadila/goproj/pkg/reader/toml"
)

const (
	ConfigName   = "goproj.conf.toml"
	configFolder = ".config/goproj/"
)

var (
	UserHomeDir, _ = os.UserHomeDir()
	ConfigPath     = filepath.Join(UserHomeDir, configFolder)
	ConfigFilepath = filepath.Join(ConfigPath, ConfigName)
)

func filepathExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func createConfigFoldersIfNotExists(configPath string) error {
	if filepathExists(configPath) {
		return nil
	}
	return folders.Create([]*models.Folder{
		models.NewFolder(configPath),
	})
}

func Store(config *models.GlobalConfig) (err error) {
	err = createConfigFoldersIfNotExists(ConfigPath)
	if err != nil {
		return
	}

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
	if !filepathExists(ConfigFilepath) {
		return
	}
	conf, err := reader.GetGlobalConfig(ConfigFilepath)
	if err != nil {
		output.Fatal(err)
	}
	return conf, true
}
