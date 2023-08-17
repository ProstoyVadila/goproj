package models

import (
	"embed"
	"log"
	"os"
	"reflect"
	"strings"
)

const (
	TEMPLATE_PATH        = "./templates"
	FILE_TO_MOVE_PATH    = "./templates/files"
	LOCALRUN_RESULT_PATH = "./tests/tempFiles"

	GITIGNORE_FILE    = ".gitignore"
	DOCKERIGNORE_FILE = ".dockerignore"
	MAIN_GO_FILE      = "main.go"
	MAKEFILE          = "Makefile"
)

var foldersToGenerate = map[string]struct{}{
	"cmd":      {},
	"pkg":      {},
	"internal": {},
}

// ProjectInfo contains all information about the new project to create.
type ProjectInfo struct {
	EmbedFiles      embed.FS
	Documents       []*Document
	Folders         []*Folder
	FilesToGenerate map[string]*Document
	FilesToSkip     []string
	FoldersToSkip   []string
	PackageName     string
	Author          string
	Description     string
	Path            string
	InitGit         bool
	InitVSCode      bool
}

func (p *ProjectInfo) AddFiles(templates ...*Document) {
	p.Documents = append(p.Documents, templates...)
}

func (p *ProjectInfo) AddFolders(folders ...*Folder) {
	p.Folders = append(p.Folders, folders...)
}

// setFilesToGenerate sets FilesToGenerate field of ProjectInfo without skipped files.
func (p *ProjectInfo) setFilesToGenerate(filesToGenerate map[string]*Document) {
	for _, fileToSkip := range p.FilesToSkip {
		fileToSkip = strings.ToLower(fileToSkip)
		delete(filesToGenerate, fileToSkip)
	}
	p.FilesToGenerate = filesToGenerate
}

// setFoldersToGenerate generates and sets Folders for ProjectInfo without skipped ones.
func (p *ProjectInfo) setFoldersToGenerate(foldersToGenerate map[string]struct{}) {
	for _, folderToSkip := range p.FoldersToSkip {
		delete(foldersToGenerate, folderToSkip)
	}
	for k := range foldersToGenerate {
		p.AddFolders(NewFolder(k))
	}
}

// generateDocuments creates Documents  for ProjectInfo by filling in the setup data into the existing configuration of templates and files.
func (p *ProjectInfo) generateDocuments() {
	for _, v := range p.FilesToGenerate {
		// setting a path where file will be created
		v.Filepath = p.Path

		if v.IsTemplate {
			var valuesFromSetup []reflect.Value

			// getting setup data from ProjectInfo by field names from the configuration
			for _, fieldName := range v.DataToAdd {
				val := reflect.ValueOf(p).Elem().FieldByName(fieldName)
				valuesFromSetup = append(valuesFromSetup, val)
			}

			// getting a constructor function from the configuration
			constructor := reflect.ValueOf(v.Constructor)
			// calling the construction function with args from the setup
			resultValue := constructor.Call(valuesFromSetup)
			// setting it as the Data field of Document for filling in the appropriate template
			v.Data = resultValue[0].Interface()
		}
		p.AddFiles(v)
	}
}

// NewProjectInfo constructs a ProjectInfo
func NewProjectInfo(setup *Setup) *ProjectInfo {
	absPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectInfo := &ProjectInfo{
		Author:        setup.Author,
		PackageName:   setup.PackageName,
		FilesToSkip:   setup.FilesToSkip,
		FoldersToSkip: setup.FoldersToSkip,
		Description:   setup.Description,
		InitGit:       setup.InitGit,
		InitVSCode:    setup.InitVSCode,
		Path:          absPath,
	}

	projectInfo.setFilesToGenerate(filesToGenerate)

	projectInfo.setFoldersToGenerate(foldersToGenerate)

	projectInfo.generateDocuments()

	return projectInfo
}
