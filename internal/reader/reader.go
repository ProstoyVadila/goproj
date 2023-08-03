package reader

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ProstoyVadila/goprojtemplate/internal/models"
)

func readInput(scanner *bufio.Scanner, previousMessage string) (string, error) {
	fmt.Print(previousMessage)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

func ReadInput() (*models.ProjectInfo, error) {

	scanner := bufio.NewScanner(os.Stdin)

	author, err := readInput(scanner, "Please, enter your name: ")
	if err != nil {
		return &models.ProjectInfo{}, err
	}
	packageName, err := readInput(scanner, "Please, enter your new project (package) name: ")
	if err != nil {
		return &models.ProjectInfo{}, err
	}
	description, err := readInput(scanner, "Please, add a description to your project: ")
	if err != nil {
		return &models.ProjectInfo{}, err
	}

	return models.NewProjectInfo(author, packageName, description), nil
}
