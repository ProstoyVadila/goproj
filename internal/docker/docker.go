package docker

import (
	"fmt"
	"net/http"
	"time"
)

const TIMEOUT = 800 * time.Millisecond

var golangDockerAPIUrl = "https://hub.docker.com/v2/repositories/library/golang/tags/%s"

func ImageExists(goVersion string) (bool, error) {
	url := fmt.Sprintf(golangDockerAPIUrl, goVersion)

	client := http.Client{
		Timeout: TIMEOUT,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}
