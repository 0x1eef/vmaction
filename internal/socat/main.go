package socat

import (
	"fmt"
	"os/exec"

	//"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Run() error {
	fmt.Println("VM doesn't have IP yet, please wait up to 100 seconds")
	cmd, err := exec.Command("bash", "socat.sh").CombinedOutput()
	if err != nil {
		return err
	}
	println(string(cmd))
	return nil
}
