package models

type ConfigFromFile struct {
	Skip        []string `yaml:"skip" json:"skip" toml:"skip"`
	Author      string   `yaml:"author" json:"author" toml:"author"`
	Description string   `yaml:"description" json:"description" toml:"description"`
	InitGit     bool     `yaml:"git" json:"git" toml:"git"`
	InitVSCode  bool     `yaml:"vscode" json:"vscode" toml:"vscode"`
}
