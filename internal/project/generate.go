package project

import (
	"embed"

	"github.com/ProstoyVadila/goproj/internal/git"
	"github.com/ProstoyVadila/goproj/internal/info.go"
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/vscode"
	"github.com/ProstoyVadila/goproj/pkg/files"
	"github.com/ProstoyVadila/goproj/pkg/folders"
	"github.com/ProstoyVadila/goproj/pkg/output"
)

// some wierd Go embed magic here
//
//go:embed templates/* templates/files/*
var EmbedFiles embed.FS

// Generate creates files and folders, initializes git repo, opens VS Code according to GlobalConfig, CLI args or Input Setup.
func Generate(ArgsSetup ...*models.Setup) {
	output.Info(info.LOGO)

	// trying to get setup from the configuration file or CLI args, or Input
	setup := enrichSetup(ArgsSetup...)

	// Show final Setup
	setup.Show()

	// aggregating all info about the projct to generate
	projectInfo := models.NewProjectInfo(setup)
	projectInfo.EmbedFiles = EmbedFiles

	// generating main folder
	output.Info("Generating project folder...")
	if err := folders.CreateOne(projectInfo.GetMainFolder()); err != nil {
		output.Err(err, "oh no")
		output.Fatal(err)
	}

	// generating files
	output.Info("Generating files...")
	if err := files.Generate(projectInfo); err != nil {
		output.Fatal(err)
	}

	// generating folders
	output.Info("Generating folders...")
	if err := folders.Create(projectInfo.Folders); err != nil {
		output.Fatal(err)
	}

	// Git init
	if projectInfo.InitGit {
		if err := git.InitGitRepo(projectInfo); err != nil {
			output.Fatal(err)
		}
	}

	output.Info("\nNew project successfully generated!")

	// open VS Code
	if projectInfo.InitVSCode {
		if err := vscode.InitVSCode(projectInfo.MainFolder); err != nil {
			output.Err(err, "cannot open VS Code")
		}
	} else {
		output.Info(
			"Just jump into %s folder: `cd /%s`\n",
			projectInfo.MainFolder,
			projectInfo.MainFolder,
		)

	}
}
