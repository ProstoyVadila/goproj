package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/ProstoyVadila/goprojtemplate/internal/git"
	"github.com/ProstoyVadila/goprojtemplate/internal/reader"
	"github.com/ProstoyVadila/goprojtemplate/pkg/files"
	"github.com/ProstoyVadila/goprojtemplate/pkg/folders"
)

//go:embed templates/* templates/files/*
var EmbedFiles embed.FS

func main() {
	fmt.Println("Let's start!")

	projectInfo, err := reader.ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	projectInfo.EmbedFiles = EmbedFiles

	err = files.Generate(projectInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = folders.Create(projectInfo.Folders)
	if err != nil {
		log.Fatal(err)
	}

	if projectInfo.InitGit {
		err = git.InitGitRepo(projectInfo)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully generated!")
}
