package models

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
	// TODO add images checks
	goVersion := GoVersion() + "-alpine"
	return &DockerfileInfo{
		GoVersion:     goVersion,
		AlpineVersion: ALPINE_VERSION,
	}
}