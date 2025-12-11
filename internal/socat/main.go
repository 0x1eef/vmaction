package socat

import (
	"fmt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
	"os/exec"
)

func Run() error {
	fmt.Println("VM doesn't have IP yet, please wait up to 200 seconds")
	if err := cmd.Run(exec.Command("bash", "socat.sh")); err != nil {
		return err
	}
	//go func() {
	//	dest := fmt.Sprintf("TCP:%s:22", string(cmd))
	//	exec.Command("socat", "TCP-LISTEN:2222,fork,reuseaddr", dest).Run()
	//}()
	return nil
}
