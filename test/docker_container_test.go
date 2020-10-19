package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/logger"
)

type StopOptions struct {
	// Seconds to wait for stop before killing the container (default 10)
	Time int

	// Set a logger that should be used. See the logger package for more info.
	Logger *logger.Logger
}

func TestDockerContainerEndpoint(t *testing.T) {
	tag := "testnode/terratest"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../code/", buildOptions)
	opts := &docker.RunOptions{
		Detach:       true,
		OtherOptions: []string{"-p=8081:8080"},
	}
	containerId := docker.Run(t, tag, opts)
	url := "http://localhost:8081"
	retries := 10
	sleep := 3 * time.Second
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		url,
		nil,
		retries,
		sleep,
		func(statusCode int, body string) bool {
			isOk := statusCode == 200
			isNginx := strings.Contains(body, "#Happy coding#")
			return isOk && isNginx
		},
	)
	fmt.Println(containerId)
	StopOptions := &docker.StopOptions{
		Time: 1,
	}
	defer docker.Stop(t, []string{containerId}, StopOptions)

}
