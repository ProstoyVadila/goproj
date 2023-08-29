package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReadmeInfo(t *testing.T) {
	author := "Bob"
	description := "bla bla"
	readme1 := &ReadmeInfo{
		AuthorName:  author,
		Description: description,
	}
	readme2 := NewReadmeInfo(author, description)
	assert.Equal(t, readme1, readme2)
}
