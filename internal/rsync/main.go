package rsync

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	dest := fmt.Sprintf("runner@%s:%s/", ip, filepath.Dir(todir))
	args := []string{"-rvah", "--mkpath", "-e", "ssh -o StrictHostKeyChecking=no", fromdir, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
