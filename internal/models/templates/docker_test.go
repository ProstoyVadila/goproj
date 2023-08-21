package templates

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DockerfileInfo(t *testing.T) {
	version := GoVersion(runtime.Version()) + "-alpine"
	dockerfile1 := &DockerfileInfo{
		GoVersion:     version,
		AlpineVersion: ALPINE_VERSION,
	}
	dockerfile2 := NewDockerfileInfo()
	assert.Equal(t, dockerfile1, dockerfile2)
}
