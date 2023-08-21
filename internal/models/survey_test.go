package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestSurvey() *Survey {
	packageName := "github.com/Alice/example_project"
	author := "Alice"
	description := "exmaple description"
	initGit := false
	initVSCode := false
	skip := []string{"Makefile", "pkg/"}

	return &Survey{
		PackageName: packageName,
		Author:      author,
		Description: description,
		Skip:        skip,
		InitGit:     initGit,
		InitVSCode:  initVSCode,
	}
}

func TestToSetup(t *testing.T) {
	survey := getTestSurvey()
	setup1 := NewSetup(
		survey.PackageName,
		survey.Author,
		survey.Description,
		survey.Skip,
		survey.InitGit,
		survey.InitVSCode,
	)
	assert.Equal(t, setup1, survey.ToSetup())
}

func TestToGlobalConfig(t *testing.T) {
	survey := getTestSurvey()
	conf1 := NewGlobalConfig(
		survey.Author,
		survey.Skip,
		survey.InitGit,
		survey.InitVSCode,
	)
	assert.Equal(t, conf1, survey.ToGlobalConfig())
}
