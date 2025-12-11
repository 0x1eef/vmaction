package main

import (
	"fmt"
	"os"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/apt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/ssh"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/vm"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/xz"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
)

func main() {
	var (
		ip      string
		archive string
		image   string
		session *ssh.Session
		err     error
	)
	if err := apt.Run(); err != nil {
		abort("error: %s\n", err)
	}
	if archive, err = curl.Run(); err != nil {
		abort("error: %s\n", err)
	}
	if image, err = xz.Run(archive); err != nil {
		abort("error: %s\n", err)
	}
	if ip, err = vm.Run(image); err != nil {
		abort("error: %s\n", err)
	}
	if session, err = ssh.Run(ip); err != nil {
		abort("error: %s\n", err)
	}
	fmt.Println("SSH session established")
	defer session.Close()
	if out, err := session.CombinedOutput(input.Run); err != nil {
		abort("error: %s\n", err)
	} else {
		fmt.Println(string(out))
	}
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}
