package models

import (
	"testing"

	"github.com/ProstoyVadila/goproj/internal/models/templates"
	"github.com/stretchr/testify/assert"
)

func getEmptyTestProjectInfo() *ProjectInfo {
	return &ProjectInfo{}
}

func TestAddFiles(t *testing.T) {
	doc := getTestDocument()
	assert.NotEmpty(t, doc)

	proj := getEmptyTestProjectInfo()
	proj.AddDocuments(doc)

	assert.NotEmpty(t, proj.Documents)
	assert.Equal(t, doc, proj.Documents[0])
}

func TestAddFolders(t *testing.T) {
	folder := NewFolder("test_folder")

	proj := getEmptyTestProjectInfo()
	proj.AddFolders(folder)

	assert.NotEmpty(t, proj.Folders)
	assert.Equal(t, folder, proj.Folders[0])
}

func TestSetFilesToGenerate(t *testing.T) {
	constructor := templates.NewReadmeInfo
	testFilesToGenerate := map[string]*Document{
		"readme": NewDocument(
			"readme.tmpl",
			"README.md",
			"readme_template_path",
			constructor,
			true,
			[]string{"Author", "Description"},
		),
		"dockerfile": NewDocument(
			"dockerfile.tmpl",
			"Dockerfile",
			"dockerfile_template_path",
			constructor,
			true,
			make([]string, 0),
		),
	}
	filesToSkip := []string{"Dockerfile"}
	doc := testFilesToGenerate["readme"]

	proj := getEmptyTestProjectInfo()
	proj.FilesToSkip = filesToSkip
	proj.setFilesToGenerate(testFilesToGenerate)

	assert.NotEmpty(t, proj.FilesToGenerate)
	assert.Len(t, proj.FilesToGenerate, 1)

	_, ok := proj.FilesToGenerate["dockerfile"]
	assert.False(t, ok)

	readme, ok := proj.FilesToGenerate["readme"]
	assert.True(t, ok)

	assert.Equal(t, doc.Name, readme.Name)
	assert.Equal(t, doc.Filename, readme.Filename)
	assert.Equal(t, doc.Filepath, readme.Filepath)
	assert.Equal(t, doc.TemplatePath, readme.TemplatePath)
	assert.Equal(t, doc.IsTemplate, readme.IsTemplate)
	assert.Equal(t, doc.DataToAdd, readme.DataToAdd)
	assert.Equal(t, doc.Data, readme.Data)

}

func TestSetFoldersToGenerate(t *testing.T) {
	testFoldersToGenerate := map[string]struct{}{
		"cmd":      {},
		"pkg":      {},
		"internal": {},
		"config":   {},
		"tests":    {},
	}
	foldersToSkip := []string{"cmd", "pkg"}

	proj := getEmptyTestProjectInfo()
	proj.FoldersToSkip = foldersToSkip

	proj.setFoldersToGenerate(testFoldersToGenerate)

	var folders []*Folder
	for folder := range testFoldersToGenerate {
		for _, skippedFolder := range foldersToSkip {
			if folder == skippedFolder {
				continue
			}
		}
		folders = append(folders, NewFolder(folder))
	}

	assert.NotEmpty(t, proj.FoldersToSkip)
	assert.NotEmpty(t, proj.Folders)
	assert.Equal(t, folders, proj.Folders)
}

func TestSetDocuments(t *testing.T) {
	constructor := templates.NewReadmeInfo
	authorName := "Alice"
	description := "test description"
	testFilesToGenerate := map[string]*Document{
		"readme": NewDocument(
			"readme.tmpl",
			"README.md",
			"readme_template_path",
			constructor,
			true,
			[]string{"Author", "Description"},
		),
	}

	doc := testFilesToGenerate["readme"]
	doc.Data = constructor(authorName, description)

	proj := getEmptyTestProjectInfo()
	proj.FilesToGenerate = testFilesToGenerate
	proj.Author = authorName
	proj.Description = description

	proj.setDocuments()

	assert.NotEmpty(t, proj.Documents)
	assert.Len(t, proj.Documents, 1)
	assert.Equal(t, doc.Name, proj.Documents[0].Name)
	assert.Equal(t, doc.Filename, proj.Documents[0].Filename)
	assert.Equal(t, doc.Filepath, proj.Documents[0].Filepath)
	assert.Equal(t, doc.TemplatePath, proj.Documents[0].TemplatePath)
	assert.Equal(t, doc.IsTemplate, proj.Documents[0].IsTemplate)
	assert.Equal(t, doc.DataToAdd, proj.Documents[0].DataToAdd)
	assert.Equal(t, doc.Data, proj.Documents[0].Data)
}

func TestNewProjectInfo(t *testing.T) {
	setup := getTestSetup(true)

	proj := NewProjectInfo(setup)

	assert.NotEmpty(t, proj)
	assert.Equal(t, setup.Author, proj.Author)
	assert.Equal(t, setup.PackageName, proj.PackageName)
	assert.Equal(t, setup.Description, proj.Description)
	assert.Equal(t, setup.InitGit, proj.InitGit)
	assert.Equal(t, setup.InitVSCode, proj.InitVSCode)
	assert.Equal(t, setup.FilesToSkip(), proj.FilesToSkip)
	assert.Equal(t, setup.FoldersToSkip(), proj.FoldersToSkip)
}
