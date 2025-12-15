package main

import (
	"fmt"
	"os"
	"path"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/apt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/keys"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/rsync"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/ssh"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/vm"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/xz"
)

func main() {
	var (
		ip, archive, image, dir, script string
		payload                         []byte
		ok                              bool
		session                         *ssh.Session
		err                             error
	)
	group("Environment", func() {
		if dir, ok = os.LookupEnv("GITHUB_WORKSPACE"); !ok {
			abort("GITHUB_WORKSPACE not set\nEnvironment: %v", os.Environ())
		}
		script = path.Join(dir, "hardenedbsd-vm.sh")
		payload = fmt.Appendf(payload, "#!/bin/sh\nset -ex\ncd %s\n%s\n", dir, input.Run)
	})
	group("Install tools", func() {
		if err := apt.Run(); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("Tools installed")
	})
	group("Download VM", func() {
		if archive, err = curl.Run(input.Release); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("VM downloaded:", archive)
	})
	group("Extract VM", func() {
		if image, err = xz.Run(archive); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("VM extracted:", image)
	})
	group("Install SSH keys", func() {
		if err := keys.Install(); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("SSH keys installed")
	})
	group("Run VM", func() {
		if ip, err = vm.Run(image); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Printf("\nPlease wait for SSH...\n")
		if session, err = ssh.Run(ip); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("SSH session established")
	})
	group("Save payload", func() {
		err = os.WriteFile(script, payload, 0755)
		if err != nil {
			abort("error: %s\n", err)
		} else {
			fmt.Printf("User input saved as %s\n", path.Base(script))
		}
	})
	group("Copy payload to VM", func() {
		if err := rsync.CopyToVM(ip, dir); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("Payload copied to VM")
	})
	group("Run payload", func() {
		defer session.Close()
		fmt.Printf("Payload: %s\n", script)
		if out, err := session.CombinedOutput(script); err != nil {
			abort("error: \n%s%s\n\n", string(out), err)
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
