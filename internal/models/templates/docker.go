package templates

import (
	"fmt"
	"runtime"

	"github.com/ProstoyVadila/goproj/internal/docker"
	"github.com/ProstoyVadila/goproj/pkg/output"
)

const (
	DOCKERFILE_TEMPLATE = "Dockerfile.tmpl"
	DOCKERFILE          = "Dockerfile"

	LATEST = "latest"
)

type DockerfileInfo struct {
	GoImageTag     string
	AlpineImageTag string
}

func NewDockerfileInfo() *DockerfileInfo {
	goVersion := GoVersion(runtime.Version())
	goImageTag := goVersion + "-alpine"
	exists, err := docker.ImageExists(goImageTag)
	if err != nil {
		msg := fmt.Sprintf("cannot check golang:%s docker image existance due %v", goImageTag, err)
		output.Warn(msg)
	}
	if !exists {
		msg := fmt.Sprintf("cannot find golang:%s docker image. Adding golang:%s to the Dockerfile instead", goImageTag, LATEST)
		output.Warn(msg)
		goImageTag = LATEST
	}
	return &DockerfileInfo{
		GoImageTag:     goImageTag,
		AlpineImageTag: LATEST,
	}
}
