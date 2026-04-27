package docker_test

import (
	"fmt"
	"testing"

	"github.com/6dainas/6dainas-app-management/pkg/docker"
)

func TestGetDir(t *testing.T) {
	fmt.Println(docker.GetDir("", "config"))
}
