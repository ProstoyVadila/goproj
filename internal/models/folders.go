package models

import (
	"io/fs"
)

const FILEMODE fs.FileMode = 0755

type Folder struct {
	Name        string
	Permissions fs.FileMode
}

func NewFolder(name string, perm ...uint32) *Folder {
	return &Folder{
		Name:        name,
		Permissions: FILEMODE,
	}
}
