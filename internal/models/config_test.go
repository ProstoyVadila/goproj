package models

import (
	"fmt"
	"testing"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/stretchr/testify/assert"
)

func getTestGlobalConfig() *GlobalConfig {
	author := "Alice"
	skip := []string{"Makefile", "pkg/"}
	initGit := true
	initVSCode := true
	prefix := "github.com/alice"

	return &GlobalConfig{
		Author:     author,
		Prefix:     prefix,
		Skip:       skip,
		InitGit:    initGit,
		InitVSCode: initVSCode,
	}
}

func Test_NewGlobalConfig(t *testing.T) {
	conf1 := getTestGlobalConfig()
	conf2 := NewGlobalConfig(
		conf1.Author,
		conf1.Prefix,
		conf1.Skip,
		conf1.InitGit,
		conf1.InitVSCode,
	)

	assert.Equal(t, conf1, conf2)
}

func TestGlobalConfig_getShow(t *testing.T) {
	conf := getTestGlobalConfig()

	omap := orderedmap.NewOrderedMap[string, any]()

	omap.Set("Author: %s", conf.Author)
	omap.Set("Global prefix: %s", conf.Prefix)
	omap.Set("Objects to skip: %v", conf.Skip)
	omap.Set("Init Git Repo: %v", conf.InitGit)
	omap.Set("Open in VS Code: %v", conf.InitVSCode)

	omapShow, msg := conf.getShow()

	assert.NotEmpty(t, msg)
	assert.NotEmpty(t, omapShow)
	assert.Equal(t, omap, omapShow)
}

func TestGlobalConfig_ShowString(t *testing.T) {
	conf := getTestGlobalConfig()
	_, msg := conf.getShow()

	showStr := conf.ShowString()

	assert.NotEmpty(t, showStr)
	assert.Contains(t, showStr, msg)
	assert.Contains(t, showStr, conf.Author)
	assert.Contains(t, showStr, fmt.Sprint(conf.Skip))
	assert.Contains(t, showStr, fmt.Sprint(conf.InitGit))
	assert.Contains(t, showStr, fmt.Sprint(conf.InitVSCode))
}

func ExampleGlobalConfig_Show() {
	conf := getTestGlobalConfig()
	conf.Show()
}
