package templates

import "runtime"

const (
	DOCKERFILE_TEMPLATE = "Dockerfile.tmpl"
	DOCKERFILE          = "Dockerfile"

	ALPINE_VERSION = "latest"
)

type DockerfileInfo struct {
	GoVersion     string
	AlpineVersion string
}

func NewDockerfileInfo() *DockerfileInfo {
	goVersion := GoVersion(runtime.Version()) + "-alpine"
	return &DockerfileInfo{
		GoVersion:     goVersion,
		AlpineVersion: ALPINE_VERSION,
	}
}
