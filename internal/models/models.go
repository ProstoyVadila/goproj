package models

import (
	"embed"
	"log"
	"os"

	"github.com/ProstoyVadila/goproj/internal/models/templates"
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

var LIST_OF_FOLDERS = [4]string{"cmd", "pkg", "internal", "tests"}

type ProjectInfo struct {
	EmbedFiles  embed.FS
	Templates   []*Document
	Folders     []*Folder
	FilesToSkip []string
	AuthorName  string
	PackageName string
	Path        string
	InitGit     bool
}

func (p *ProjectInfo) AddFiles(templates ...*Document) {
	p.Templates = append(p.Templates, templates...)
}

func (p *ProjectInfo) AddFolders(folders ...*Folder) {
	p.Folders = append(p.Folders, folders...)
}

func NewProjectInfo(setup *Setup) *ProjectInfo {
	absPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectInfo := &ProjectInfo{
		AuthorName:  setup.Author,
		PackageName: setup.PackageName,
		FilesToSkip: setup.FilesToSkip,
		Path:        absPath,
		// TODO get InitGit value from input questions
		InitGit: true,
	}

	// Template files
	license := NewDocument(
		templates.LICENSE_TEMPLATE,
		templates.LICENSE_FILE,
		absPath,
		TEMPLATE_PATH,
		true,
		templates.NewLicenceInfo(setup.Author),
	)
	readme := NewDocument(
		templates.README_TEMPLATE,
		templates.README_FILE,
		absPath,
		TEMPLATE_PATH,
		true,
		templates.NewReadmeInfo(setup.Author, setup.Description),
	)
	gomod := NewDocument(
		templates.GOMOD_TEMPLATE,
		templates.GOMOD_FILE,
		absPath,
		TEMPLATE_PATH,
		true,
		templates.NewGoModInfo(setup.PackageName),
	)
	dockerfile := NewDocument(
		templates.DOCKERFILE_TEMPLATE,
		templates.DOCKERFILE,
		absPath,
		TEMPLATE_PATH,
		true,
		templates.NewDockerfileInfo(),
	)

	// Files to copy
	gitignore := NewDocument(
		GITIGNORE_FILE,
		GITIGNORE_FILE,
		absPath,
		FILE_TO_MOVE_PATH,
		false,
		struct{}{},
	)
	dockerignore := NewDocument(
		DOCKERIGNORE_FILE,
		DOCKERIGNORE_FILE,
		absPath,
		FILE_TO_MOVE_PATH,
		false,
		struct{}{},
	)
	mainGoFile := NewDocument(
		MAIN_GO_FILE,
		MAIN_GO_FILE,
		absPath,
		FILE_TO_MOVE_PATH,
		false,
		struct{}{},
	)
	makefile := NewDocument(
		MAKEFILE,
		MAKEFILE,
		absPath,
		FILE_TO_MOVE_PATH,
		false,
		struct{}{},
	)

	projectInfo.AddFiles(
		license,
		readme,
		gomod,
		dockerfile,
		gitignore,
		mainGoFile,
		dockerignore,
		makefile,
	)

	for _, folderName := range LIST_OF_FOLDERS {
		projectInfo.AddFolders(NewFolder(folderName))
	}
	return projectInfo
}
