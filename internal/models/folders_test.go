package models

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFolder(t *testing.T) {
	var perm uint32 = 0755
	name := "pkg"

	folder1 := &Folder{
		Name:        name,
		Permissions: fs.FileMode(perm),
	}
	folder2 := NewFolder(name, perm)
	folder3 := NewFolder(name)

	assert.Equal(t, folder1, folder2)
	assert.Equal(t, folder1, folder3)
}
