package templates

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
	goVersion := GoVersion() + "-alpine"
	return &DockerfileInfo{
		GoVersion:     goVersion,
		AlpineVersion: ALPINE_VERSION,
	}
}
