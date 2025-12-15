package rsync

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func CopyToVM(ip, dir string) error {
	dest := fmt.Sprintf("runner@%s:%s/", ip, filepath.Dir(dir))
	args := []string{"-rvah", "--mkpath", "-e", "ssh -o StrictHostKeyChecking=no", dir, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
