package ssh

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Run(ip string) error {
	args := []string{
		"-p", "",
		"ssh",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
		"-p", "2222",
		fmt.Sprintf("root@%s", ip),
		"true",
	}
	attempts := 0
	max := 100
	for {
		if err := cmd.Run(exec.Command("sshpass", args...)); err != nil {
			attempts++
			if attempts >= max {
				return err
			}
			time.Sleep(1 * time.Second)
			fmt.Println("")
		} else {
			break
		}
	}
	return nil
}
