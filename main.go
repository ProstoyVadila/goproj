// package groprojscript
package main

import (
	"fmt"
	"log"

	"github.com/ProstoyVadila/goprojtemplate/internal/files"
	"github.com/ProstoyVadila/goprojtemplate/internal/reader"
)

func main() {
	fmt.Println("Let's start!")
	projectInfo, err := reader.ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	err = files.Generate2(projectInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully generated!")
}
