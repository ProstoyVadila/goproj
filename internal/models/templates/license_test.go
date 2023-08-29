package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLicenseInfo(t *testing.T) {
	author := "Alice"
	license1 := &LicenseInfo{
		AuthorName: author,
		Year:       2023,
	}
	license2 := NewLicenceInfo(author)
	assert.Equal(t, license1, license2)
}
