// package groprojscript
package main

import (
	"fmt"
	"log"

	"github.com/ProstoyVadila/goprojtemplate/internal/reader"
	"github.com/ProstoyVadila/goprojtemplate/pkg/files"
)

func main() {
	fmt.Println("Let's start!")
	projectInfo, err := reader.ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	err = files.Generate(projectInfo)
	if err != nil {
		log.Fatal(err)
	}
}
