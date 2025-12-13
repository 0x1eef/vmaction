package rsync

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func CopyToVM(ip string) error {
	var (
		wrkdir string
		ok     bool
	)
	if wrkdir, ok = os.LookupEnv("GITHUB_WORKSPACE"); !ok {
		return fmt.Errorf("GITHUB_WORKSPACE not set\nEnvironment: %v", os.Environ())
	}
	src := wrkdir
	dest := fmt.Sprintf("runner@%s:~/", ip)
	args := []string{"-rvah", "-e", "ssh -o StrictHostKeyChecking=no", src, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
