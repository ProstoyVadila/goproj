package models

import (
	"fmt"
	"testing"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/stretchr/testify/assert"
)

func getTestSetup(isDefault bool) *Setup {
	if isDefault {
		return &Setup{
			PackageName: "example_project",
			Author:      "Alice",
			Description: "example description",
			Skip:        []string{"Makefile", "pkg/"},
			InitGit:     true,
			InitVSCode:  false,
		}
	}
	return NewSetup(
		"test_update",
		"test_Bob",
		"test_description",
		[]string{"Makefile", "Dockerfile", "pkg/"},
		false,
		true,
	)
}

func Test_NewSetup(t *testing.T) {
	setup1 := getTestSetup(true)
	setup2 := NewSetup(
		setup1.PackageName,
		setup1.Author,
		setup1.Description,
		setup1.Skip,
		setup1.InitGit,
		setup1.InitVSCode,
	)

	assert.Equal(t, setup1, setup2)
}

func Test_NewSetupFromConfig(t *testing.T) {
	setup1 := getTestSetup(true)

	conf := NewGlobalConfig(
		setup1.Author,
		setup1.Skip,
		setup1.InitGit,
		setup1.InitVSCode,
	)
	setup2 := NewSetupFromConfig(conf)

	assert.NotEqual(t, setup1, setup2)
	assert.Equal(t, setup1.Author, setup2.Author)
	assert.Equal(t, setup1.Skip, setup2.Skip)
	assert.Equal(t, setup1.InitGit, setup2.InitGit)
	assert.Equal(t, setup1.InitVSCode, setup2.InitVSCode)
}

func Test_FilesToSkip(t *testing.T) {
	setup := getTestSetup(true)

	filesToSkip := []string{"Makefile", "Dockerfile"}
	skip := []string{"pkg/", "internal/"}
	skip = append(skip, filesToSkip...)

	setup.Skip = skip

	assert.Equal(t, filesToSkip, setup.FilesToSkip())
}

func Test_FoldersToSkip(t *testing.T) {
	setup := getTestSetup(true)

	foldersToSkip := []string{"pkg", "internal"}
	skip := []string{"Makefile", "Dockerfile"}
	for _, folder := range foldersToSkip {
		skip = append(skip, folder+"/")
	}

	setup.Skip = skip

	assert.Equal(t, foldersToSkip, setup.FoldersToSkip())
}

func Test_Update(t *testing.T) {
	originalSetup := getTestSetup(true)

	testCases := []struct {
		name     string
		toUpdate *Setup
		another  *Setup
		original *Setup
		testFunc func(t *testing.T, toUpdate, another, original *Setup)
	}{
		{
			name:     "case 1: full update",
			toUpdate: getTestSetup(true),
			another:  getTestSetup(false),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEqual(t, original, toUpdate)
				assert.Equal(t, another, toUpdate)
			},
		},
		{
			name:     "case 2: empty author",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				"",
				originalSetup.Description,
				originalSetup.Skip,
				originalSetup.InitGit,
				originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEqual(t, another, toUpdate)
				assert.NotEmpty(t, toUpdate.Author)
				assert.Equal(t, original.Author, toUpdate.Author)

			},
		},
		{
			name:     "case 3: empty description",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				"",
				originalSetup.Skip,
				originalSetup.InitGit,
				originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEqual(t, another, toUpdate)
				assert.NotEmpty(t, toUpdate.Description)
				assert.Equal(t, original.Description, toUpdate.Description)
			},
		},
		{
			name:     "case 4: not empty skip",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				originalSetup.Description,
				[]string{"main.go", "cmd/"},
				originalSetup.InitGit,
				originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEqual(t, original, toUpdate)
				assert.NotEmpty(t, toUpdate.Skip)
				assert.Equal(t, another.Skip, toUpdate.Skip)
			},
		},
		{
			name:     "case 5: empty skip",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				originalSetup.Description,
				make([]string, 0),
				originalSetup.InitGit,
				originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEmpty(t, toUpdate.Skip)
				assert.Equal(t, original.Skip, toUpdate.Skip)
			},
		},
		{
			name:     "case 6: init git true",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				originalSetup.Description,
				originalSetup.Skip,
				true,
				originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.True(t, toUpdate.InitGit)
			},
		},
		{
			name:     "case 7: init git not equal original",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				originalSetup.Description,
				originalSetup.Skip,
				!originalSetup.InitGit,
				originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEqual(t, original.InitGit, toUpdate.InitGit)
				assert.Equal(t, another.InitGit, toUpdate.InitGit)
			},
		},
		{
			name:     "case 8: init vs code true",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				originalSetup.Description,
				originalSetup.Skip,
				originalSetup.InitGit,
				true,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.True(t, toUpdate.InitVSCode)
			},
		},
		{
			name:     "case 9: init vs code not equal original",
			toUpdate: getTestSetup(true),
			another: NewSetup(
				originalSetup.PackageName,
				originalSetup.Author,
				originalSetup.Description,
				originalSetup.Skip,
				originalSetup.InitGit,
				!originalSetup.InitVSCode,
			),
			original: originalSetup,
			testFunc: func(t *testing.T, toUpdate, another, original *Setup) {
				assert.Equal(t, original, toUpdate)
				toUpdate.Update(another)

				assert.NotEqual(t, original.InitVSCode, toUpdate.InitVSCode)
				assert.Equal(t, another.InitVSCode, toUpdate.InitVSCode)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.testFunc(t, tt.toUpdate, tt.another, tt.original)
		})
	}
}

func TestSetup_getShow(t *testing.T) {
	setup := getTestSetup(true)

	omap := orderedmap.NewOrderedMap[string, any]()

	omap.Set("Project (package) name: %s", setup.PackageName)
	omap.Set("Author: %s", setup.Author)
	omap.Set("Description: %s", setup.Description)
	omap.Set("Files to skip: %v", setup.FilesToSkip())
	omap.Set("Folders to skip: %v", setup.FoldersToSkip())
	omap.Set("Init Git Repo: %v", setup.InitGit)
	omap.Set("Open in VS Code: %v", setup.InitVSCode)

	omapShow, msg := setup.getShow()

	assert.NotEmpty(t, msg)
	assert.NotEmpty(t, omapShow)
	assert.Equal(t, omap, omapShow)
}

func TestSetup_ShowString(t *testing.T) {
	setup := getTestSetup(true)
	_, msg := setup.getShow()

	showStr := setup.ShowString()

	assert.NotEmpty(t, showStr)
	assert.Contains(t, showStr, msg)
	assert.Contains(t, showStr, setup.Author)
	assert.Contains(t, showStr, setup.Description)
	assert.Contains(t, showStr, setup.PackageName)
	assert.Contains(t, showStr, fmt.Sprint(setup.FilesToSkip()))
	assert.Contains(t, showStr, fmt.Sprint(setup.FoldersToSkip()))
	assert.Contains(t, showStr, fmt.Sprint(setup.InitGit))
	assert.Contains(t, showStr, fmt.Sprint(setup.InitVSCode))
}

func ExampleSetup_Show() {
	setup := getTestSetup(true)
	setup.Show()
}
