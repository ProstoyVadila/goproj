package templates

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerfileInfo(t *testing.T) {
	version := GoVersion(runtime.Version()) + "-alpine"
	dockerfile1 := &DockerfileInfo{
		GoImageTag:     version,
		AlpineImageTag: LATEST,
	}
	dockerfile2 := NewDockerfileInfo()
	assert.Equal(t, dockerfile1, dockerfile2)
}
