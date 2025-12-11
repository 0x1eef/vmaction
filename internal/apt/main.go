package apt

import (
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

var (
	packages = []string{
		"zstd",
		"libvirt-daemon-system",
		"virt-manager",
		"qemu-kvm",
		"libosinfo-bin",
		"sshpass",
		"socat",
	}
)

func Run() error {
	commands := [][]string{
		{"sudo", "apt-get", "update"},
		append([]string{"sudo", "apt-get", "install", "-y"}, packages...),
	}
	for _, c := range commands {
		err := cmd.Run(exec.Command(c[0], c[1:]...))
		if err != nil {
			return err
		}
	}
	return nil
}
