package templates

import "time"

const (
	LICENSE_TEMPLATE = "LICENSE.tmpl"
	LICENSE_FILE     = "LICENSE"
)

type LicenseInfo struct {
	Version     string
	AuthorName  string
	LicenseFile string
	Year        int
}

func NewLicenceInfo(authorName string) *LicenseInfo {
	return &LicenseInfo{
		AuthorName: authorName,
		Year:       time.Now().Year(),
	}
}
