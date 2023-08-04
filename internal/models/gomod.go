package models

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
