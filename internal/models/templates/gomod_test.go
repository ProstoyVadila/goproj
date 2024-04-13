package templates

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoVersion(t *testing.T) {
	testCases := []struct {
		name    string
		version string
		result  string
	}{
		{
			name:    "case 1",
			version: "go1.20.6",
			result:  "1.20",
		},
		{
			name:    "case 2",
			version: "mem4.20.0",
			result:  "4.20",
		},
		{
			name:    "case 3",
			version: "go3.12",
			result:  "3.12",
		},
		{
			name:    "case 4",
			version: "go1-20",
			result:  "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GoVersion(tt.version), tt.result)
		})
	}

}

func TestNewGoModInfo(t *testing.T) {
	packageName := "new_project"
	gomod1 := &GoModInfo{
		PackageName: packageName,
		GoVersion:   GoVersion(runtime.Version()),
	}
	gomod2 := NewGoModInfo(packageName)
	assert.Equal(t, gomod1, gomod2)
}
