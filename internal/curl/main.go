package curl

import (
	"fmt"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
)

var (
	Destination = "image.raw"
	URLMap      = map[string]string{
		"16-CURRENT": "FIXME",
		"15-STABLE":  "https://github.com/0x1eef/hardenedbsd-builder/actions/runs/20037038355/artifacts/4800738241",
		"14-STABLE":  "FIXME",
	}
)

func Run() (string, error) {
	if downloadURL, err := url(); err != nil {
		return "", err
	} else {
		args := []string{"-L", "-o", Destination, downloadURL}
		return Destination, cmd.Run(exec.Command("curl", args...))
	}
}

func url() (string, error) {
	u, ok := URLMap[input.Release]
	if !ok {
		return "", fmt.Errorf("unknown release: %s", input.Release)
	}
	return u, nil
}
