package main

import (
	"fmt"
	"os"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/apt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/ssh"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/vm"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/xz"
)

func main() {
	var (
		ip      string
		archive string
		image   string
		session *ssh.Session
		err     error
	)
	group("Install tools", func() {
		if err := apt.Run(); err != nil {
			abort("error: %s\n", err)
		}
	})
	group("Download VM", func() {
		if archive, err = curl.Run(); err != nil {
			abort("error: %s\n", err)
		}
	})
	group("Extract VM", func() {
		if image, err = xz.Run(archive); err != nil {
			abort("error: %s\n", err)
		}
	})
	group("Run VM", func() {
		if ip, err = vm.Run(image); err != nil {
			abort("error: %s\n", err)
		}
		if session, err = ssh.Run(ip); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("SSH session established")
	})
	group("Run input script", func() {
		defer session.Close()
		if out, err := session.CombinedOutput(input.Run); err != nil {
			abort("error: %s\n", err)
		} else {
			fmt.Println(string(out))
		}
	})
}

func group(label string, fn func()) {
	fmt.Printf("::group::%s\n", label)
	fn()
	fmt.Println("::endgroup::")
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}
