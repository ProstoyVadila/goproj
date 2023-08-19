package models

import "fmt"

var showString = `
You set this configuration for your new projects:

Author: %s
Description: %s
Obejcts to Skip: %v
Init Git Repo: %v
Open in VS Code: %v

`

type GlobalConfig struct {
	Skip        []string `yaml:"skip" json:"skip" toml:"skip"`
	Author      string   `yaml:"author" json:"author" toml:"author"`
	Description string   `yaml:"description" json:"description" toml:"description"`
	InitGit     bool     `yaml:"git" json:"git" toml:"git"`
	InitVSCode  bool     `yaml:"vscode" json:"vscode" toml:"vscode"`
}

func (g GlobalConfig) Show() {
	fmt.Println(g.ShowString())
}

func (g GlobalConfig) ShowString() string {
	return fmt.Sprintf(
		showString,
		g.Author,
		g.Description,
		g.Skip,
		g.InitGit,
		g.InitVSCode,
	)
}
