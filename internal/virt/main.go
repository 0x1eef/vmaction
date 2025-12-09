package virt

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Run(image string) error {
	vm := "testvm"
	if err := create(vm, image); err != nil {
		return err
	}
	if err := wait(); err != nil {
		return err
	}
	return nil
}

func create(vm, image string) error {
	args := []string{
		"--name", vm,
		"--memory", "6144",
		"--vcpus", "2",
		"--arch", "x86_64",
		"--disk", "path="+image + ",format=raw,bus=virtio",
		"--os-variant", "freebsd13.1",
		"--graphics", "none",
		"--network", "none",
		"--qemu-commandline='-netdev user,id=net0,hostfwd=tcp::2222-:22 -device virtio-net-pci,netdev=net0'",
		"--noautoconsole",
		"--import",
		"--serial", "file,path=/tmp/testvm.log",
	}
	return cmd.Run(exec.Command("virt-install", args...))
}

func wait() error {
	maxAttempts := 100
	attempts := 0
	for {
		time.Sleep(2 * time.Second)
		if out, err := exec.Command(
			"sshpass", "-p", "",
			"ssh", "-p", "2222", "root@localhost",
			"-o ConnectTimeout=2", "-o ConnectionAttempts=1",
			"true",
		).CombinedOutput(); err != nil {
			attempts++
			fmt.Println(string(out))
			fmt.Printf("VM not online yet (attempt %d of %d)\n", attempts, maxAttempts)
			if attempts == maxAttempts {
				return err
			}
		} else {
			fmt.Printf("VM is online\n")
			break
		}
	}
	return nil
}
