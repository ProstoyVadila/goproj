package models

import (
	"strings"

	"github.com/ProstoyVadila/goproj/pkg/output"
	"github.com/elliotchance/orderedmap/v2"
)

type GlobalConfig struct {
	Skip       []string `yaml:"skip" json:"skip" toml:"skip"`
	Author     string   `yaml:"author" json:"author" toml:"author"`
	Prefix     string   `yaml:"prefix" json:"prefix" toml:"prefix"`
	InitGit    bool     `yaml:"git" json:"git" toml:"git"`
	InitVSCode bool     `yaml:"vscode" json:"vscode" toml:"vscode"`
}

func NewGlobalConfig(author, prefix string, skip []string, initGit, initVSCode bool) *GlobalConfig {
	return &GlobalConfig{
		Author:     author,
		Skip:       skip,
		InitGit:    initGit,
		InitVSCode: initVSCode,
		Prefix:     prefix,
	}
}

func (g *GlobalConfig) ValidatePrefix() {
	githubPrefix := "github.com/"
	if !strings.Contains(g.Prefix, githubPrefix) {
		g.Prefix = githubPrefix + g.Prefix + "/"
	}
	g.Prefix = strings.ReplaceAll(g.Prefix, "//", "/")
}

// getShow creates ordered map of GlobalConfig fields and msg for output.
func (g *GlobalConfig) getShow() (*orderedmap.OrderedMap[string, any], string) {
	msg := "This is your global config:"
	omap := orderedmap.NewOrderedMap[string, any]()

	omap.Set("Author: %s", g.Author)
	omap.Set("Global prefix: %s", g.Prefix)
	omap.Set("Objects to skip: %v", g.Skip)
	omap.Set("Init Git Repo: %v", g.InitGit)
	omap.Set("Open in VS Code: %v", g.InitVSCode)

	return omap, msg
}

// Show writes GlobalConfig info to standart output.
func (g *GlobalConfig) Show() {
	output.Show(g.getShow())
}

// ShowString returns output string for GlobalConfig.
func (g *GlobalConfig) ShowString() string {
	return output.ShowString(g.getShow())
}
