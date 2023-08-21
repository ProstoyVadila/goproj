package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGlobalConfig(t *testing.T) {
	author := "Alice"
	skip := []string{"Makefile", "pkg/"}
	initGit := true
	initVSCode := true

	conf1 := &GlobalConfig{
		Author:     author,
		Skip:       skip,
		InitGit:    initGit,
		InitVSCode: initVSCode,
	}
	conf2 := NewGlobalConfig(
		author,
		skip,
		initGit,
		initVSCode,
	)

	assert.Equal(t, conf1, conf2)
}
