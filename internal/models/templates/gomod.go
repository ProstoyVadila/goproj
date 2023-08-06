package templates

import (
	"regexp"
	"runtime"
)

const (
	GOMOD_TEMPLATE = "go.mod.tmpl"
	GOMOD_FILE     = "go.mod"
)

type GoModInfo struct {
	GoVersion   string
	PackageName string
}

func NewGoModInfo(packageName string) *GoModInfo {
	return &GoModInfo{
		PackageName: packageName,
		GoVersion:   GoVersion(),
	}
}

// GoVersion gets the version of Go from the runtime in the format <major_version>.<minor_version> (1.19 for example).
func GoVersion() string {
	fullVersion := runtime.Version()
	re := regexp.MustCompile(`\d\.\d+`)
	return re.FindString(fullVersion)
}
