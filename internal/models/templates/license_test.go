package templates

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLicenseInfo(t *testing.T) {
	author := "Alice"
	license1 := &LicenseInfo{
		AuthorName: author,
		Year:       time.Now().Year(),
	}
	license2 := NewLicenceInfo(author)
	assert.Equal(t, license1, license2)
}
