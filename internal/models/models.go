package models

import (
	"embed"
	"log"
	"os"
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

// TODO get values from input/setup
var LIST_OF_FOLDERS = [4]string{"cmd", "pkg", "internal", "tests"}

type ProjectInfo struct {
	Templates   []*Document
	EmbedFiles  embed.FS
	AuthorName  string
	PackageName string
	Path        string
	InitGit     bool
	Folders     []*Folder
}

func (p *ProjectInfo) AddFiles(templates ...*Document) {
	p.Templates = append(p.Templates, templates...)
}

func (p *ProjectInfo) AddFolders(folders ...*Folder) {
	p.Folders = append(p.Folders, folders...)
}

func NewProjectInfo(authorName, packageName, description string) *ProjectInfo {
	absPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectInfo := &ProjectInfo{
		AuthorName:  authorName,
		PackageName: packageName,
		Path:        absPath,
		// TODO get InitGit value from input questions
		InitGit: true,
	}

	// Template files
	license := NewDocument(
		LICENSE_TEMPLATE,
		LICENSE_FILE,
		absPath,
		TEMPLATE_PATH,
		true,
		NewLicenceInfo(authorName),
	)
	readme := NewDocument(
		README_TEMPLATE,
		README_FILE,
		absPath,
		TEMPLATE_PATH,
		true,
		NewReadmeInfo(authorName, description),
	)
	gomod := NewDocument(
		GOMOD_TEMPLATE,
		GOMOD_FILE,
		absPath,
		TEMPLATE_PATH,
		true,
		NewGoModInfo(packageName),
	)
	dockerfile := NewDocument(
		DOCKERFILE_TEMPLATE,
		DOCKERFILE,
		absPath,
		TEMPLATE_PATH,
		true,
		NewDockerfileInfo(),
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
