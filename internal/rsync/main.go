package rsync

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func CopyToVM(ip string) error {
	var (
		fromdir string
		todir   string
		err     error
		ok      bool
	)
	if fromdir, err = os.Getwd(); err != nil {
		return err
	}
	if todir, ok = os.LookupEnv("GITHUB_WORKSPACE"); !ok {
		return fmt.Errorf("GITHUB_WORKSPACE not set\nEnvironment: %v", os.Environ())
	}
	dest := fmt.Sprintf("runner@%s:%s/", ip, todir)
	args := []string{"-rvah", "--mkpath", "-e", "ssh -o StrictHostKeyChecking=no", fromdir, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
