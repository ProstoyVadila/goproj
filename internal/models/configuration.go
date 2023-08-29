package models

import "github.com/ProstoyVadila/goproj/internal/models/templates"

const (
	TEMPLATE_PATH        = "./templates"
	FILE_TO_MOVE_PATH    = "./templates/files"
	LOCALRUN_RESULT_PATH = "./tests/tempFiles"

	GITIGNORE_FILE    = ".gitignore"
	DOCKERIGNORE_FILE = ".dockerignore"
	MAIN_GO_FILE      = "main.go"
	MAKEFILE          = "Makefile"
	ENVFILE           = ".env"
)

// Configuration fo each file
var filesToGenerate = map[string]*Document{
	// Templates
	"license": NewDocument(
		templates.LICENSE_TEMPLATE,
		templates.LICENSE_FILE,
		TEMPLATE_PATH,
		templates.NewLicenceInfo,
		true,
		[]string{"Author"},
	),
	"readme": NewDocument(
		templates.README_TEMPLATE,
		templates.README_FILE,
		TEMPLATE_PATH,
		templates.NewReadmeInfo,
		true,
		[]string{"Author", "Description"},
	),
	"go.mod": NewDocument(
		templates.GOMOD_TEMPLATE,
		templates.GOMOD_FILE,
		TEMPLATE_PATH,
		templates.NewGoModInfo,
		true,
		[]string{"PackageName"},
	),
	"dockerfile": NewDocument(
		templates.DOCKERFILE_TEMPLATE,
		templates.DOCKERFILE,
		TEMPLATE_PATH,
		templates.NewDockerfileInfo,
		true,
		[]string{},
	),
	// Files to copy
	".gitignore": NewDocument(
		GITIGNORE_FILE,
		GITIGNORE_FILE,
		FILE_TO_MOVE_PATH,
		struct{}{},
		false,
		[]string{},
	),
	".dockerignore": NewDocument(
		DOCKERIGNORE_FILE,
		DOCKERIGNORE_FILE,
		FILE_TO_MOVE_PATH,
		struct{}{},
		false,
		[]string{},
	),
	"main.go": NewDocument(
		MAIN_GO_FILE,
		MAIN_GO_FILE,
		FILE_TO_MOVE_PATH,
		struct{}{},
		false,
		[]string{},
	),
	"makefile": NewDocument(
		MAKEFILE,
		MAKEFILE,
		FILE_TO_MOVE_PATH,
		struct{}{},
		false,
		[]string{},
	),
	".env": NewDocument(
		ENVFILE,
		ENVFILE,
		FILE_TO_MOVE_PATH,
		struct{}{},
		false,
		[]string{},
	),
}
