package models

import (
	"embed"
	"fmt"
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

var foldersToGenerate = [4]string{"cmd", "pkg", "internal", "tests"}

// ProjectInfo contains all information about the new project to create.
type ProjectInfo struct {
	EmbedFiles      embed.FS
	Documents       []*Document
	Folders         []*Folder
	FilesToGenerate map[string]*Document
	FilesToSkip     []string
	PackageName     string
	Author          string
	Description     string
	Path            string
	SkipGit         bool
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

// setDocuments sets Documents field of ProjectInfo by filling in the setup data to the existing configuration of templates and files.
func (p *ProjectInfo) setDocuments() {
	for _, v := range p.FilesToGenerate {
		// setting a path where file will be created
		v.Filepath = p.Path

		if v.IsTemplate {
			var valuesFromSetup []reflect.Value

			// getting data from setup by field names from the configuration
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
	fmt.Println("Abs Path", absPath)

	projectInfo := &ProjectInfo{
		Author:      setup.Author,
		PackageName: setup.PackageName,
		FilesToSkip: setup.FilesToSkip,
		Description: setup.Description,
		Path:        absPath,
		// TODO get InitGit value from input questions
		SkipGit: false,
	}
	projectInfo.setFilesToGenerate(filesToGenerate)
	projectInfo.setDocuments()

	for _, folderName := range foldersToGenerate {
		projectInfo.AddFolders(NewFolder(folderName))
	}
	return projectInfo
}
